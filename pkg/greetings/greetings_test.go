package greetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	greeting := Greet()

	// Check that greeting is not empty
	assert.NotEmpty(t, greeting)

	// Check that greeting is one of the expected values
	expectedGreetings := []string{
		"Hello, World!",
		"Welcome to edays-golang!",
		"Greetings from Go!",
		"Hello there!",
		"Good day!",
		"Nice to meet you!",
	}

	assert.Contains(t, expectedGreetings, greeting)
}

func TestGreetPerson(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid name",
			input:    "John",
			expected: "Hello, John! Welcome to edays-golang!",
		},
		{
			name:     "empty name",
			input:    "",
			expected: "Hello, anonymous!",
		},
		{
			name:     "name with spaces",
			input:    "John Doe",
			expected: "Hello, John Doe! Welcome to edays-golang!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GreetPerson(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGreetWithTime(t *testing.T) {
	// We can't easily mock time in this simple implementation,
	// so we'll just test that the function returns a greeting with time
	greeting := GreetWithTime()

	// Check that greeting contains "Current time:"
	assert.Contains(t, greeting, "Current time:")

	// Check that greeting is not empty
	assert.NotEmpty(t, greeting)
}

func TestGreetRandomness(t *testing.T) {
	// Test that Greet() can return different values
	greetings := make(map[string]int)

	// Run multiple times to check for randomness
	for i := 0; i < 100; i++ {
		greeting := Greet()
		greetings[greeting]++
	}

	// We should have multiple different greetings
	assert.Greater(t, len(greetings), 1, "Greet() should return different greetings")
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet()
	}
}

func BenchmarkGreetPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreetPerson("John")
	}
}
