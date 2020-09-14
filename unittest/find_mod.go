package unittest

import (
	"fmt"
	"math"
)

func findMod(input float64) string {
	if input == 0 {
		return fmt.Sprintf("%.2f : all", input)
	} else if math.Mod(input, 15) == 0 {
		return fmt.Sprintf("%.2f : FizzBuzz", input)
	} else if math.Mod(input, 5) == 0 {
		return fmt.Sprintf("%.2f : Buzz", input)
	} else if math.Mod(input, 3) == 0 {
		return fmt.Sprintf("%.2f : Fizz", input)
	} else {
		return fmt.Sprintf("%.2f", input)
	}

}
