package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"strconv"
	"strings"
	"sync"
)

func fillDummyIndex(es *elasticsearch.Client, count int) {
	var wg sync.WaitGroup
	for i := 1; i < count; i += 1 {
		wg.Add(1)
		go func (n int){
			res, err := es.Index(
				"test2",
				strings.NewReader(`{"title": "test docs, will be visible to search ` + strconv.Itoa(n) + `"}`),
				es.Index.WithRefresh("true"),
			)
			if err != nil {
				log.Fatalf("ERROR: %s", err)
			}
			defer res.Body.Close()

			log.Println(res)
			wg.Done()
			fmt.Println("es inserted", n)
		}(i)
	}
	wg.Wait()
	fmt.Println("es inserted all docs")
}
