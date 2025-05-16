package main

import "fmt"

func GoFunction1() {
	fmt.Println("Good Evening")
}

//Function with parameters
func GoFunction2(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("Sum is", sum)
}

//Function with multiple return values
func GoFunction3(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

//Function with multiple return values
func GoFunction4(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
