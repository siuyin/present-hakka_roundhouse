package main

import "fmt"

func main() {
	fmt.Println("a traditional monolith")
	fmt.Printf("2+3 = %v\n", sum(2, 3))
}

func sum(a, b int) int {
	return a + b
}
