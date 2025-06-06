package main

import (
	"fmt"
	"time"
)

func main() {
	// สร้าง Channel สำหรับส่งข้อมูลประเภท string
	ch := make(chan string)
	// เรียกใช้ Goroutine เพื่อส่งข้อมูลไปยัง Channel
	go func() {
		time.Sleep(1 * time.Second) // จำลองการทำงานที่ใช้เวลา
		ch <- "Hello from Goroutine!"

	}()
	// รับข้อมูลจาก Channel
	msg := <-ch
	fmt.Println(msg) // ผลลัพธ์: Hello from Goroutine!
}
