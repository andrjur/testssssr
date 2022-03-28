package main

import (
	"fmt"
)

func ExampleScope2() {
	var i string = "Строка"
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)
	for i := -1; i < 1; i++ {
		fmt.Println("before truth", i)
		i := true
		fmt.Println("after truth", i)
	}
	fmt.Println(i)
}

func main() {
	ExampleScope2()
}
