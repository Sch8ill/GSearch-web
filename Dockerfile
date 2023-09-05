FROM golang:1.20-bullseye

WORKDIR /gsearch-web

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o gsearch-web

EXPOSE 8000

CMD ["./gsearch-web"]