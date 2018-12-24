package main

import (
	// "context"
	// "fmt"
	// "log"

	// elastic "github.com/olivere/elastic"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

// Human structure to serialize/deserialize data in Elasticsearch
type Human struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"human":{
			"properties":{
				"Name":{
					"type":"text"
				},
				"Gender":{
					"type":"text"
				},
				"Age":{
					"type":"text"
				}
			}
		}
	}
}`

func main() {
	// index means to deposit a new document into elastic search
	// context	// lets you define how much you want input into the terminal
	ctx := context.Background()

	// Creates a new client server
	eclient, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	// Pings the server to see if its running
	info, code, err := eclient.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Server running \n")
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// check if people indexExist
	// Use the IndexExists service to check if a specified index exists.
	exists, err := eclient.IndexExists("people").Do(ctx)
	if err != nil {
		// log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		panic(err)
	}
	if !exists {
		// create a new index
		createIndex, err := eclient.CreateIndex("people").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Printf("Hello Sir")
			//Not Acknowledged
		}
	}

	/*
		Indexing people to database (using JSON serialization):
			* Salman
			* Kevin
	*/

	salman := Human{Name: "Salman", Gender: "Male", Age: "20"}
	set1, err := eclient.Index(). // slappin in the salman index into the eclient
					Index("people"). // refer to the people dic that was created
					Type("human").   // refer to the struct type that you made to serialize/deserialize
					Id("1").
					BodyJson(salman).
					Do(ctx)

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	} else {
		log.Println("Success!")
	}
	fmt.Printf("Indexed human %s to index %s, type %s\n", set1.Id, set1.Index, set1.Type)

	kevin := Human{Name: "Kevin", Gender: "Male", Age: "19"}
	set2, err := eclient.Index(). // slappin in the kevin index into the eclient
					Index("people"). // refer to the people dic that was created
					Type("human").   // refer to the struct type that you made to serialize/deserialize
					Id("2").
					BodyJson(kevin).
		// BodyString(set2). // referring to set2 body string of kevin
		Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	} else {
		log.Println("Success!")
	}
	fmt.Printf("Indexed human %s to index %s, type %s\n", set2.Id, set2.Index, set2.Type)

	/*
		Getting people from the database (using JSON serialization):
			* Salman
			* Kevin
	*/

	// Get tweet with specified ID
	get1, err := eclient.Get().
		Index("people").
		Type("human").
		Id("1").
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

	// Searching with a term Query
	termQuery := elastic.NewTermQuery("name", "salman") // this term looks for Salman
	searchResult, err := eclient.Search().              // Search through dic and error if not there
								Index("people").
								Query(termQuery).
								From(0).Size(2).
								Pretty(true). // pretty prints the values
								Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	// fmt.Printf("Found a total of %d people\n", searchResult.TotalHits())

	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d people\n", searchResult.Hits.TotalHits)
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index
			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Human
			err := json.Unmarshal(*hit.Source, &t) // decoding hit to a Human variable
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			fmt.Printf("People:\n")
			fmt.Printf("%s %s %s\n", t.Name, t.Gender, t.Age)
		}
	} else {
		// No hits
		fmt.Print("Found no People\n")
	}

	// fmt.Println(searchResult)
	// var ttyp Person
	// for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
	// 	if t, ok := item.(Person); ok {
	// 		fmt.Printf("Person by %s: %s\n", t.User, t.Message)
	// 	}
	// }

}
