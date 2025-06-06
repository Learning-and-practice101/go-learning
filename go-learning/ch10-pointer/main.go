package main

import "fmt"

func main() {
	SamplePointer1()
	fmt.Println("===========")

	age := 25

	//แสดงอายุเติม
	fmt.Println("อายุเดิม:", age)

	ChageAge(&age)

	SamplePointer2()
}
