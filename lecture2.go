package main

import "fmt"

// Multi value return
func multi(a, b int) (int, int) {
	return a + b, a * b
}

func myfunc() {
	i := 0
	i++
	fmt.Println(i)
}

// int in between refers to the return type
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

	arr := [3]int{1, 2, 3}
	// when := is used, you dont need the var before the variable and type after
	fmt.Println(arr[0])
	splice1 := arr[0:3]
	// splice1.append(5)
	fmt.Println(len(splice1))

	// fmt.Println(len(splice1))
	// var var1 int = 3
	// fmt.Println(var1)
	// a := [3] {1, 2, 3} defines int array with 3 elements
	numbers := make(map[string]int)

	numbers["one"] = 1 // assign value by key
	numbers["ten"] = 10
	numbers["three"] = 3
	fmt.Println("The third number is: ", numbers["three"]) // get values
	// It prints: The third number is: 3
	x := 7
	if x > 10 {
		fmt.Printf("yoyoyo")
	} else {
		fmt.Printf("fuck")
	}
	myfunc()
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println(sum)
	// Or use for loops to act like a while loop
	// break is used the same way
	sum1 := 1
	for sum1 < 100 {
		sum1 += sum1
	}
	fmt.Println(sum1)
	a := 10
	b := 20
	fmt.Println(max(a, b))
	fmt.Println(multi(a, b))

}
