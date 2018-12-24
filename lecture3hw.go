package main

import (
	"fmt"
	"os"
	"strconv"
)

// Error handling
// var user = os.Getenv("USER")

// func init() {
// 	if user == "" {
// 		panic("no value for $USER")
// 	}
// }
	/*
	// Structs
	type Product struct {
		name          string
		itemID        int
		cost          float32
		isAvailable   bool
		inventoryLeft int
	}
	// Initializing struct
	goBook := Product{}
	goBook.name = "Webapps in Go"
	fmt.Printf("The product's name is %s\n", goBook.name)

	// List of structs
	const (
		WHITE = iota
        BLACK
        BLUE
        RED
		YELLOW
	)
	type Color byte
	type Box struct { 
		width, height, depth float64
		color Color
	}
	type BoxList []Box 
	func boxVolumes(b Box) Volume() float64 {
		return b.width * b.height * b.depth
	}
	*/

type Stringer interface { 
	String() string
}

type Human struct {
	name	string
	age		int
	phone	string
}

//Human implements fmt.Stringer
func (h Human) String() string { 
	return "Name:" + h.name + ", Age:" + 
		strconv.Itoa(h.age) + " years, Contact:" +\ h.phone
}

// embedded interfaces 
type

func main() {
	Bob := Human{"Bob", 39, "000-2222-XXX"}
	fmt.Println("This human is: ", Bob)
}

