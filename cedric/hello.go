package main

import (
	"fmt"
)

func main() {
	var s = make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("emp:", s)
	var slice = s[0:1]
	fmt.Println("slice:", slice)
	slice[0] = "x"
	fmt.Println("slice:", s)
}
