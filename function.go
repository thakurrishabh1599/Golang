package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous function")
	}()
	val := func() int {
		return 10
	}
	fmt.Println(val())
	func(v ...string) {
		fmt.Println(v)
	}("test", "best")

}

// 	a := a()
// 	b := b()
// 	// sum(a, b)
// 	// passByRef(&a, &b)
// 	variadicfunc(a, b, 11)

// }
// func a() int {
// 	return 10
// }

// func b() int {
// 	return 10
// }
// func variadicfunc(v ...int) {
// 	fmt.Println(v)
// 	sum := 0
// 	for _, num := range v {
// 		sum += num
// 	}
// 	fmt.Println(sum)
// }

// // func sum(a, b int) {
// // 	fmt.Println(a + b)
// // }
// // func passByRef(a, b *int) {
// // 	fmt.Println(*a + *b)
// // }
