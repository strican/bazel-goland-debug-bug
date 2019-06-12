package main

import "fmt"

func Foo(a, b int) int {
	return a + b
}

func main() {
	fmt.Print(Foo(5, 6))
}
