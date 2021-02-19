package main

import "fmt"

// func main() {
// 	fmt.Println("counting")

// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i)
// 	}

// 	defer fmt.Println("last defer")

// 	fmt.Println("done")
// }

func test1() (myInt int) {
	myInt = 1

	defer func() {
		myInt++
	}()

	return myInt
}

func main() {
	fmt.Printf("test1: %d\n", test1())
}
