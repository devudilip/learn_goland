package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		}
	}
}

func foo() {
	fmt.Println("i am in foo")
}
