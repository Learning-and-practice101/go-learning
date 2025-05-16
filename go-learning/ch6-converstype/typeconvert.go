// แปลงจำนวนเต็มเป็นจำนวนทศนิยม
package main

import (
	"fmt"
	"strconv"
)

// int to float
func TypeConvert1() {

	var i int = 42
	var f float64 = float64(i) // แปลง int เป็น float64
	fmt.Printf("i = %d, f = %f\n", i, f)
}

// แปลงจากจำนวนทศนิยมเป็นจำนวนเต็ม
func TypeConvert2() {
	var f float64 = 42.56
	var i int = int(f) // แปลง float64 เป็น int
	fmt.Printf("f = %f, i = %d\n", f, i)
}

// แปลงจากจำนวนเต็มเป็นสตริง
func TypeConvert3() {
	var i int = 42
	var s string = strconv.Itoa(i) // Interget to Ascii
	fmt.Printf("i = %d, s = %s\n", i, s)
}

// แปลงจากสตริงเป็นจำนวนเต็ม
func TypeConvert4() {
	var s string = "42x"
	var i int
	var err error
	i, err = strconv.Atoi(s) // Ascii to Interget
	// ถ้าแปลงไม่สำเร็จ err จะไม่เป็น nil
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("s = %s, i = %d\n", s, i)
}

// แปลงจากบูลีนเป็นจำนวนเต็ม (แสดงในรูปแบบ integer เพราะ Go ไม่รองรับการแปลงจาก bool เป็น int โดยตรง)
func TypeConvert5() {
	var b bool = true
	var i int
	if b {
		i = 1 // แปลง true เป็น 1
	} else {
		i = 0 // แปลง false เป็น 0
	}
	fmt.Printf("b = %t, i = %d\n", b, i)
}
