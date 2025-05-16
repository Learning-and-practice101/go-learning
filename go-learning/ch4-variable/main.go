package main

import (
	"fmt"
)

// การประกาศตัวแปรใน Go
// แบบ Global variable
var xmlglobal string = "outside main"

func main() {

	fmt.Println(xmlglobal) // outside main

	// แบบ Local variable
	xmllocal := "inside main"
	fmt.Println(xmllocal) // inside main

	//a lot for variable
	var (
		num   float32
		count int    = 10
		name  string = "john"
	)

	num = 1.5

	fmt.Println(num)
	fmt.Println(count)
	fmt.Println(name)

	//Multiple short variable declaration
	// _ คั้นระหว่างตัวเลข ไม่สามารถ run , ได้
	var x, y, z = 10_000_000, 2_0 + 5, 30
	var p, q, r = "Hello", "world", "Go"

	fmt.Println(x, y, z) //10 20 30
	fmt.Println(p, q, r) //Hello world go

	// ประกาศค่าคงที่
	// const <constant_name> <data_type> = <value>
	const pi float32 = 3.14
	// pi = 3.14159 // Error: cannot assign to pi
	const vat = 7.0
	const name1, name2 = "John", "Doe"

	fmt.Println(pi)    // 3.14
	fmt.Println(vat)   // 7
	fmt.Println(name1) // John
}
