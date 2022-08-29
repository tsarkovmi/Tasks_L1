package main

import (
	"fmt"
)

func main() {
	var number int64
	var bit, index int

	number = 63
	bit = 0
	index = 2

	var mask int64 = 1 << (index - 1)
	if bit == 1 {
		fmt.Printf("           mask: %08b\n", mask)
		fmt.Printf("original number: %08b\n", number)
		fmt.Printf("  result number: %08b\n", number^mask)
		return
	}

	fmt.Printf("           mask: %08b\n", mask)
	fmt.Printf("original number: %08b\n", number)
	fmt.Printf("  result number: %08b\n", number&^mask)
}
