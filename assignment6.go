package main

// Relationship
// MySQL => Databases => Tables => Columns/Rows
// Elasticsearch => Indices => Types => Documents with Properties

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

// "context"
// "fmt"
// "log"

// elastic "github.com/olivere/elastic"

// "context"
// "fmt"
// "log"

// elastic "github.com/olivere/elastic"

type Owner struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Pet struct {
	Name         string `json:"name"`
	Owner        string `json:"OwnerDocID"`
	Breed        string `json:"breed"`
	Microchipped bool   `json:"microchipped"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 2,
		"number_of_replicas": 0
	},
	"mappings":{
		"owner":{
			"properties":{
				"Name":{
					"type":"text"
				},
				"Location":{
					"type":"text"
				}
			}
		},
		"pet":{
			"properties":{
				"Name":{
					"type":"text"
				},
				"Owner":{
					"type":"number"
				},
				"Breed":{
					"type":"text"
				},
				"Microchipped":{
					"type":"bolean"
				}
			}
		}
	}
}`

func main() {
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

	// check if owner & pet indexExist
	// Use the IndexExists service to check if a specified index exists.
	exists, err := eclient.IndexExists("owner").Do(ctx)
	if err != nil {
		// log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		panic(err)
	}
	if !exists {
		// create a new index
		createOwner, err := eclient.CreateIndex("owners").BodyString(mapping).Do(ctx) //
		createPet, err := eclient.CreateIndex("pets").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createOwner.Acknowledged || !createPet.Acknowledged {
			fmt.Printf("Hello Sir")
			//Not Acknowledged
		}
	}

	salman := Owner{Name: "Salman", Location: "Atlantic Avenue"}
	set1, err := eclient.Index().
		Index("salman"). // refer to the people dic that was created
		Type("owner").   // refer to the struct type that you made to serialize/deserialize
		Id("1").
		BodyJson(salman).
		Do(ctx)

	if err != nil {
		// log.SetFlags(log.lstdFlags | log.Lshortfile)
		log.Println(err)
	} else {
		log.Println("Success!")
	}

	get1, err := eclient.Get().
		Index("owners").
		Type("owner").
		Id("1").
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	doggy := Pet{Name: "Scooby Doo", Owner: get1.Id, Breed: "Husky", Microchipped: true}
	set2, err := eclient.Index().
		Index("pets"). // refer to the people dic that was created
		Type("pet").   // refer to the struct type that you made to serialize/deserialize
		Id("1").
		BodyJson(doggy).
		Do(ctx)

	if err != nil {
		// log.SetFlags(log.lstdFlags | log.Lshortfile)
		log.Println(err)
	} else {
		log.Println("Success!")
	}

	fmt.Printf(doggy.Owner)

	fmt.Printf("Slapped owner %s to %s database, of type %s\n", set1.Id, set1.Index, set1.Type)
	fmt.Printf("Slapped pet %s to %s database, of type %s\n", set2.Id, set2.Index, set2.Type)
}
