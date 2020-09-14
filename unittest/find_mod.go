package unittest

import (
	"fmt"
	"math"
)

func findMod(input int) string {
	n := float64(input)
	if n == 0 {
		return fmt.Sprint(input)
	} else if math.Mod(n, 15) == 0 {
		return mod15(input)

		// return fmt.Sprint(input, " : FizzBuzz")
	} else if math.Mod(n, 5) == 0 {
		return fmt.Sprint(input, " : Buzz")
	} else if math.Mod(n, 3) == 0 {
		// return fmt.Sprint(input, " : Fizz")
		return "Fizz"
	} else {
		return fmt.Sprintf("%.0f", n)
	}

}
func mod15(input int) string {
	n := float64(input)

	if math.Mod(n, 15) == 0 {
		return "FizzBuzz"
	}
	return fmt.Sprintf("%d", input)

}
