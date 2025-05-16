package main

import (
	"fmt"
	_ "math"

	"github.com/google/uuid"
	"github.com/thanaschai/greeting"
	// personal "github.com/thanaschai/greeting"
)

func main() {
	fmt.Println("Welcome to go")

	var x int = 5
	_ = x
	//generate a new UUID
	//สร้างตัวแปรเพื่อเก็บ UUID
	var id string = uuid.New().String()
	//print UUID
	fmt.Println("UUID:", id)

	greeting.SayGreeting()
	//greeting.SayMeeting()

	//personal.PrintPersonal
}
