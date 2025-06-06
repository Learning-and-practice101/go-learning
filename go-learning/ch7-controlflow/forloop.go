package main

import (
	"fmt"
)

func Forloop1() {

	for i := 1; i <= 5; i++ {
		fmt.Println("Value of i is", i)
	}
}

// For like a while loop
func Forloop2() {
	i := 1
	for i <= 5 {
		fmt.Println("Value of i is", i)
		i++
	}
}

// For infinite loop
func Forloop3() {

	i := 1

	for {
		fmt.Println("Infinite loop", i)
		i++
	}

}

// For infinite loop
func Forloop4() {

	i := 1

	for {

		if i == 1024 {
			break
			// continue
		}

		fmt.Println("Infinite loop", i)

		i++
	}
}
