package main

import (
	"fmt"
)


func main() {
	var i interface{} = 1.0

	s := i.(string)
	fmt.Println(s)

	// s, ok := i.(string)
	// fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
