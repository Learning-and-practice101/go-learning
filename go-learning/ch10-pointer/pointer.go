package main

import (
	"fmt"
)

func SamplePointer1() {
	// บ้านเลขที่ 10 (เก็บค่าจริง)
	value := 10

	// ป้ายเลขที่บ้าน = ตัวแปร pointer ที่เก็บ address ของบ้าน
	pointer := &value

	// แสดง address ของบ้าน
	fmt.Println("บ้านเลขที่ (ที่อยู่ของ value):", pointer) // 0xc00000a0d8
	// แสดงค่าที่อยู่ในบ้าน
	fmt.Println("ค่าภายในบ้าน (ค่าของ value):", value) // 10
	// แสดงค่าภายในบ้าน โดยการแสดงค่าที่ pointer ชี้ไป
	fmt.Println("ค่าภายในบ้าน (ค่าของ pointer):", *pointer) // 10

	// เปลี่ยนค่าภายในบ้าน โดยการใช้ pointer
	*pointer = 20
	//value :=30
	//pass by referage ดีกกว่า

	// แสดงค่าภายในบ้าน
	fmt.Println("ค่าภายในบ้าน (ค่าของ value):", value) // 20

	fmt.Println("============================")

	// การใช้ pointer กับ slice
	// สร้าง slice ที่มีค่าเริ่มต้น
	slice := []int{1, 2, 3}

	// สร้าง pointer ที่ชี้ไปยัง slice
	pointerSlice := &slice

	// แสดง address ของ slice
	fmt.Println("บ้านเลขที่ (ที่อยู่ของ slice):", pointerSlice)

	// แสดงค่าภายใน slice
	fmt.Println("ค่าภายในบ้าน (ค่าของ slice):", *pointerSlice) // [1 2 3]

	// เปลี่ยนค่าภายใน slice โดยการใช้ pointer
	(*pointerSlice)[0] = 10

	// แสดงค่าภายใน slice
	fmt.Println("ค่าภายในบ้าน (ค่าของ slice):", *pointerSlice) // [10 2 3]
}

// func chageAge(age int)  {
// 	age = 20
// }
//pass by value

func ChageAge(age *int) { //get pointer
	*age += 1 //เพื่มค่าใน pointer
	fmt.Println("อายุเดิม:", *age)
}

// กรณีใช้ pointer
// กำหนด struct สำหรับห้อง
type Room struct {
	RoomNumber int
	BedColor   string
}

// ฟังก์ชันที่ใช้กุญแจ (pointer) เพื่อเปลี่ยนสีเตียง
func changeBedColorWithPointer(keyToRoom *Room, newColor string) {
	keyToRoom.BedColor = newColor // เปลี่ยนสีเตียงผ่าน pointer
}

func SamplePointer2() {
	// สร้างห้องหมายเลข 101 มีเตียงสีเหลือง
	room101 := Room{
		RoomNumber: 101,
		BedColor:   "เหลือง",
	}

	fmt.Println("ก่อนเปลี่ยนสีเตียง:")
	fmt.Println("หมายเลขห้อง:", room101.RoomNumber)
	fmt.Println("สีเตียง:", room101.BedColor) // เหลือง

	fmt.Println("----------------------")

	// เรียกใช้ฟังก์ชันเพื่อเปลี่ยนสีเตียง
	changeBedColorWithPointer(&room101, "เขียว") // ส่ง pointer ของห้องไปให้ฟังก์ชัน

	fmt.Println("\nหลังเปลี่ยนสีเตียง:")
	fmt.Println("หมายเลขห้อง:", room101.RoomNumber)
	fmt.Println("สีเตียง:", room101.BedColor) // เขียว
}
