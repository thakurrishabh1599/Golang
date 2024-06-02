package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 4; i++ {
	// 	fmt.Println(i)
	// }
	// //Infinite
	// // for {
	// // 	fmt.Println("Hello Friends")
	// // }
	// //While loop
	// j := 0
	// for j < 10 {
	// 	fmt.Println(j)
	// 	j++
	// }
	//Range loop
	// for i, j := range "goguruji" {
	// 	fmt.Println("The index number of %c is %d\n", j, i)
	// }
	// chnl := make(chan int)
	// go func() {
	// 	chnl <- 10
	// 	chnl <- 100
	// 	chnl <- 1000
	// 	chnl <- 10000
	// 	close(chnl)

	// }()
	// for i := range chnl {
	// 	fmt.Println(i)
	// }
	mmap := map[int]string{
		0: "a",
		1: "b",
		2: "c",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}
}
