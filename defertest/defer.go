//Package defertest は
package defertest

import "fmt"

func DeferTest() (myInt int) {
	myInt = 1

	defer func() {
		myInt++
	}()

	return myInt
}

func F() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	G(0)
	fmt.Println("Returned normally from g.")
}

func G(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	G(i + 1)
}

// func main() {
// 	fmt.Println("counting")

// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i)
// 	}

// 	defer fmt.Println("last defer")

// 	fmt.Println("done")
// }

// func main() {
// 	fmt.Printf("test1: %d\n", test1())
// }

// func main() {
// 	f()
// 	fmt.Println("Returned normally from f.")
// }
