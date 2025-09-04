package greetings

import (
	"fmt"
	"math/rand"
	"time"
)

// Greet returns a random greeting message
func Greet() string {
	greetings := []string{
		"Hello, World!",
		"Welcome to edays-golang!",
		"Greetings from Go!",
		"Hello there!",
		"Good day!",
		"Nice to meet you!",
	}

	rand.Seed(time.Now().UnixNano())
	return greetings[rand.Intn(len(greetings))]
}

// GreetPerson returns a personalized greeting
func GreetPerson(name string) string {
	if name == "" {
		return "Hello, anonymous!"
	}
	return fmt.Sprintf("Hello, %s! Welcome to edays-golang!", name)
}

// GreetWithTime returns a greeting with current time
func GreetWithTime() string {
	now := time.Now()
	greeting := Greet()
	return fmt.Sprintf("%s Current time: %s", greeting, now.Format("2006-01-02 15:04:05"))
}
