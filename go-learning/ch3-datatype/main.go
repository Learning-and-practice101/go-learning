package main

import (
	"fmt"
	"html"
)

func main() {
	//data type in Go
	//have 2 type
	//1.basic types (in,float,string,bool)
	//2.composite types (array, slice ,map ,struct,interface)

	//basic types
	// Numberic types
	var a int8 = 10   // range -128 to 127
	var b int16 = 20  // range -32768 to 32767
	var c int32 = 30  // range -2147483648 to 2147483647
	var d int64 = 40  // range -9223372036854775808 to 9223372036854775807
	var e int = 50    // range -2147483648 to 2147483647 (32-bit)
	var f uint8 = 60  // range 0 to 255
	var g uint16 = 70 // range 0 to 65535
	var h uint32 = 80 // range 0 to 4294967295
	var i uint64 = 90 // range 0 to 18446744073709551615
	var j uint = 100  // range 0 to 4294967295 (32-bit)

	// Float types
	var k float32 = 1.1 // range -3.402823466E+38 to 3.402823466E+38
	var l float64 = 2.2 // range -1.7976931348623157E+308 to 1.7976931348623157E+308

	// String type
	var m string = "Hello, World!" // range 0 to 4294967295 (32-bit)

	// Boolean type
	var n bool = true // true or false

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
	fmt.Println("g:", g)
	fmt.Println("h:", h)
	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println("j:", k)
	fmt.Println("j:", l)
	fmt.Println("j:", m)
	fmt.Println("j:", n)

	//String data type
	//Interface string คือสามารถคำนวณได้
	//can not new line
	//ใส่ตัวแปรได้
	var str1 string = "Welco`j`me to Go!\n iiiiii\\iuuuuuuuu\"frist time\""

	//Raw string --`` back tick shotcut : ถ้าจะพิมพ์ `` ให้กด alt 96--
	//can new line
	//ไม่สามารถใส่ตัวแปร
	var str2 string = `Welcome to Go!
	dfsdfs"sss"df`

	// Escape sequence (อักขระพิเศษ)
	// \n = New line (ขึ้นบรรทัดใหม่)
	// \t = Tab (tab space)
	// \\ = Backslash (เครื่องหมาย \)
	// \" = Double quote (เครื่องหมาย " )
	// ' = Single quote (เครื่องหมาย ' )
	// \r = Carriage return (กลับไปที่จุดเริ่มต้นของบรรทัด)
	// \b = Backspace (ลบอักขระก่อนหน้า)
	fmt.Println(str1)
	fmt.Println(str2)

	var str string = `<a href="https://www.google.com">Google</a>`
	// Escape HTML
	var str_esc = html.EscapeString(str)
	// Unescape HTML
	var str_unesc = html.UnescapeString(str_esc)
	fmt.Println(str_esc)
	fmt.Println(str_unesc)
}
