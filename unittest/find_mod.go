package unittest

import "fmt"

func findMod(input int) string {
	mod := [2]bool{input%5 == 0, input%3 == 0}
	switch mod {
	case [2]bool{true, true}:
		return "FizzBuzz"
	case [2]bool{true, false}:
		return "Fizz"
	case [2]bool{false, true}:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", input)

	}

}
