package greeting

import (
	"fmt"

	personal "github.com/thanaschai/greeting/internal"
)

// Public function
func SayGreeting() {
	personal.PrintPersonal()
	fmt.Println("Hello, From the greeting package!")
	sayhi()
}
