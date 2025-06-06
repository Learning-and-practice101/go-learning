package main

import (
	"fmt"
	"time"
)

func brewCoffee() {
	fmt.Println("Starting to brew coffee...")
	time.Sleep(2 * time.Second) // จำลองงานที่ใช้เวลา 2 วินาที
	fmt.Println("Coffee is ready!")
}
func bakeBread() {
	fmt.Println("Starting to bake bread...")
	time.Sleep(3 * time.Second) // จำลองงานที่ใช้เวลา 3 วินาที
	fmt.Println("Bread is ready!")
}
func washDishes() {
	fmt.Println("Starting to wash dishes...")
	time.Sleep(1 * time.Second) // จำลองงานที่ใช้เวลา 1 วินาที
	fmt.Println("Dishes are clean!")
}

func main() {
	// เริ่มต้น goroutine สำหรับการชงกาแฟ
	go brewCoffee()
	// เริ่มต้น goroutine สำหรับการอบขนมปัง
	go bakeBread()
	// เริ่มต้น goroutine สำหรับการล้างจาน
	go washDishes()
	// รอให้ goroutine ทั้งหมดทำงานเสร็จ
	time.Sleep(4 * time.Second) // รอให้ทุกงานเสร็จ
	fmt.Println("All tasks are done!")
	// หมายเหตุ: ในการใช้งานจริง ควรใช้ sync.WaitGroup แทนการใช้ time.Sleep เพื่อให้แน่ใจว่า goroutine ทั้งหมดทำงานเสร็จ
}
