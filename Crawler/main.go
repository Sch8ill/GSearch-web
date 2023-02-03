package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)


func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(res.StatusCode)
	bytes, _ :=ioutil.ReadAll(res.Body)

	fmt.Println(string(bytes))
}