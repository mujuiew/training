package unittest

import "fmt"

func findMod(input int) string {
	// n := float64(input)
	// if n == 0 {
	// 	return fmt.Sprint(input)
	// } else if math.Mod(n, 15) == 0 {
	// 	return mod15(input)

	// 	// return fmt.Sprint(input, " : FizzBuzz")
	// } else if math.Mod(n, 5) == 0 {
	// 	return fmt.Sprint(input, " : Buzz")
	// } else if math.Mod(n, 3) == 0 {
	// 	// return fmt.Sprint(input, " : Fizz")
	// 	return "Fizz"
	// } else {
	// 	return fmt.Sprintf("%.0f", n)
	// }
	mod3, mod5 := mod15(input)
	if mod3 && mod5 {
		return "FizzBuzz"
	} else if mod5 {
		return "Fizz"
	} else if mod3 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", input)
	}

}
func mod15(input int) (bool, bool) {
	return ((input % 3) == 0), ((input % 5) == 0)

}
