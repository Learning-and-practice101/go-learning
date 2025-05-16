package main

import (
	"fmt"
)

func Ifcondition() {
	number := 15
	// การใช้ if condition
	if number > 0 {
		fmt.Printf("%d Positive number\n", number)
	} else {
		fmt.Println("%d Negative number\n", number)
	}
}

func IfElseIfCondition() {
	number1 := 20
	number2 := 20

	if number1 == number2 {
		fmt.Printf("Result: %d == %d", number1, number2)
	} else if number1 > number2 {
		fmt.Printf("Result: %d > %d", number1, number2)
	} else {
		fmt.Printf("Result: %d < %d", number1, number2)
	}
}

func SwitchCase() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}

func SwitchMultiple() {

	dayOfWeek := "Wednesday"

	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	default:
		fmt.Println("Invalid day")
	}
}
