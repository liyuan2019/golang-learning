package main

import "fmt"

// func main() {
// 	var a []int // slice
// 	for i := 0; i < 5; i++ {
// 		a = append(a, i)
// 	}

// 	fmt.Printf("1. a: %v\n", a)

// 	modifySlice(a)

// 	fmt.Printf("3. a: %v\n", a)
// }

// func modifySlice(a []int) {
// 	a[2] = 9
// 	fmt.Printf("2. a: %v\n", a)
// }

func main() {
	var a []int
	for i := 0; i < 5; i++ {
		a = append(a, i)
	}

	fmt.Printf("1. a: %v\n", a)

	appendSlice(a)

	fmt.Printf("3. a: %v\n", a)
}

func appendSlice(a []int) {
	for i := 0; i < 5; i++ {
		a = append(a, i)
	}
	a[2] = 9
	fmt.Printf("2. a: %v\n", a)
}
