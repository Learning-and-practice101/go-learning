package main

import "fmt"

func Defer() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

//defer working by behigh to face
