package db

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/sch8ill/gsearch-web/config"
)

type Site struct {
	Url       string
	Text      []string
	Timestamp string
	Score     float64
}

type DBClient struct {
	uri    string
	client mongo.Client
}

const (
	DBTimeout time.Duration = 10 * time.Second
)

func New(uri string) *DBClient {
	return &DBClient{
		uri: uri,
	}
}

// connects the underlying client and tests the connection
func (dbc *DBClient) Connect() {
	log.Debug().Msg("Connecting to database...")

	ctx, _ := context.WithTimeout(context.Background(), DBTimeout)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbc.uri))
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	dbc.client = *client

	// Perform ping to test the connection
	if err := dbc.client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal().Err(err).Msg("")
	}
	log.Debug().Msg("Connected to database")
}

// closes the underlying client
func (dbc *DBClient) Close() error {
	return dbc.client.Disconnect(context.TODO())
}

// returns a database collection based on database and collection name
func (dbc *DBClient) GetColl(dbName string, collName string) *mongo.Collection {
	return dbc.client.Database(dbName).Collection(collName)
}

func (dbc *DBClient) TextSearch(query string) []Site {
	coll := dbc.client.Database(config.DBName).Collection(config.SiteColl)

	filter := bson.D{{"$text", bson.D{{"$search", query}}}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []Site
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}
