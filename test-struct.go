package main

// import (
// 	"fmt"
// )

// type st struct {
// 	id   int
// 	name string
// }

// func main() {
// 	d := st{1, "Jo"}
// 	fmt.Println(d, "値渡し前")
// 	fValue(d)
// 	fmt.Println(d, "値渡し後、outer変わらない")
// 	fPoint(&d)
// 	fmt.Println(d, "ポインター渡し後、outer変わった")
// }

// func fValue(s st) {
// 	s.id++
// 	s.name = "of"
// 	fmt.Println(s, "値渡し関数")
// }

// func fPoint(s *st) {
// 	s.id++
// 	s.name = "of"
// 	fmt.Println(*s, "ポインター渡し関数")
// }
