package main

import (
	"github.com/olivere/elastic"
)

// index function: depositing new function to elastic search
// deposit data within an index function includes
// index, type. bodyJson(test object variable created globally)
// cmd is jsut an interface 

func main() {}
	ctx := context.Background()

	eclient, err := elastic.NewClient(elastic.SetURL("https://localhost:9200"))

	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
}
