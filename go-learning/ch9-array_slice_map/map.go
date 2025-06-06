package main

import "fmt"

func SampleMap1() {
	score := make(map[string]int)
	score["Alice"] = 90
	score["Bob"] = 85
	score["Charlie"] = 95

	fmt.Println("Scores:", score)

	// การเข้าถึงสมาชิกใน Map โดยใช้ key
	fmt.Println("Score of Alice:", score["Alice"]) // 90
	fmt.Println("----------------------")
	// การตรวจสอบว่ามี key อยู่ใน Map หรือไม่
	if score, exists := score["Bob"]; exists {
		fmt.Println("Score of Bob:", score) // 85
	} else {
		fmt.Println("Bob not found")
	}

	fmt.Println("----------------------")

	// การลบสมาชิกใน Map
	delete(score, "Charlie")
	fmt.Println("Scores after deleting Charlie:", score) // Scores after deleting Charlie: map[Alice:90 Bob:85]
	fmt.Println("----------------------")
	// การวนลูปแสดงผลข้อมูลใน Map
	for name, score := range score {
		fmt.Printf("%s: %d\n", name, score)
	}

	fmt.Println("----------------------")
	// การสร้าง Map โดยใช้ literal syntax
	// โดย key เป็น string, value เป็น int
	grades := map[string]int{
		"Zooo":    100,
		"Math":    95,
		"English": 90,
		"Science": 85,
		"History": 80,
	}
	// การแสดงผลข้อมูลใน Map
	fmt.Println("Grades:", grades) // Grades: map[English:90 History:80 Math:95 Science:85]
}
