package main

import (
	"fmt"
	"log"

	"github.com/cedrichegglin/edays-golang/pkg/greetings"
)

func main() {
	// Basic greeting
	fmt.Println("Basic greeting:")
	fmt.Println(greetings.Greet())
	fmt.Println()

	// Personalized greeting
	fmt.Println("Personalized greeting:")
	fmt.Println(greetings.GreetPerson("Alice"))
	fmt.Println()

	// Greeting with time
	fmt.Println("Greeting with time:")
	fmt.Println(greetings.GreetWithTime())
	fmt.Println()

	// Multiple greetings to show randomness
	fmt.Println("Multiple greetings (showing randomness):")
	for i := 0; i < 5; i++ {
		fmt.Printf("Greeting %d: %s\n", i+1, greetings.Greet())
	}
}
