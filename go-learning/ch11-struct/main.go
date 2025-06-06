package main

import "fmt"

// กำหนด struct ชื่อ Person
type Person struct {
	Name string
	Age  int
}

// สร้าง function ชื่อ sayHello สำหรับ struct Person
func (p Person) SayHello() string {
	return "Hello " + p.Name
}

// กำหนด function ชื่อ HaveBirthday สำหรับ Person
func (p *Person) HaveBirthday() {
	p.Age++ // เพิ่มอายุ 1 ปี
}
func main() {
	// สร้าง instance ของ Person
	p := Person{
		Name: "John",
		Age:  30,
	}
	// แสดงข้อมูลของ Person
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	// เรียกฟังก์ชัน sayHello
	fmt.Println("Greeting:", p.SayHello())
	// เรียกฟังก์ชัน HaveBirthday
	p.HaveBirthday()
	fmt.Println("New Age:", p.Age)
}
