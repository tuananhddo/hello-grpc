package main

import (
	"fmt"
)

func main() {

	var a = 5
	var p = &a // copy by reference
	var q = &p

	fmt.Println("a =  ", a)   // a =  5
	fmt.Println("&a = ", &a)  // a =  5
	fmt.Println("p =  ", p)   // p =  0x10414020
	fmt.Println("*p = ", *p)  // *p =  5
	fmt.Println("&p = ", &p)  // &p =  0x1040c128
	fmt.Println("q =  ", q)   // x =  5
	fmt.Println("q =  ", **q) // x =  5

}
