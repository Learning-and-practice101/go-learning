package main

import "fmt"

func main() {
	// example
	/*
		dddd
		ddd
		dd
		d
	*/
	fmt.Println("hello world")
	fmt.Println("hello world", "hello world")
	fmt.Println("hello world" + "world")
	fmt.Println(20 + 10)

	//การ print ไม่ขึ้นบรรทัดใหม่
	fmt.Print("Hello, ")
	fmt.Print("world! ")

	//การ print format

	fmt.Printf("Hello, %s! you are %d years old.\n", "John", 30)
}
