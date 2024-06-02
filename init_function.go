package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	a := a()
	fmt.Println(a)
}
func a() int {
	return 10
}

func init() {
	fmt.Println("Init first")

}
func init() {
	fmt.Println("Init first1")

}
func init() {
	fmt.Println("Init first2")

}
func init() {
	fmt.Println("Init first3")

}
func init() {
	fmt.Println("Init first4")

}
