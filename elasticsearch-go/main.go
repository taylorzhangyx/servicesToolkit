package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil{
		log.Println("trying to connect failed")
		log.Println(err)
		return
	}
	log.Println("trying to connect pass")
	log.Println(elasticsearch.Version)
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

	fillDummyIndex(es, 100)
}
