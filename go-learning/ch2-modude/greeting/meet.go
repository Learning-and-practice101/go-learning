package greeting

import (
	"fmt"
)

// Public function
func SayMeeting() {
	fmt.Println("Hello, From the meeting package!")
}

// Private function
func sayhi() {
	fmt.Println("Hello, bro")
}
