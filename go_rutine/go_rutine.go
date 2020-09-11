package main

import (
	"fmt"
	"math"
)

// var pmt float64

func calPmt1(dis float64, i float64, num int, message chan<- float64) {
	i = i / 100 / 12
	n := float64(num)
	pmt := dis / ((1 - (1 / math.Pow((1+i), n))) / i)
	fmt.Printf("Payment Amount with Nember of payment %v = %.2f\n", num, pmt)
	message <- pmt

	fmt.Println()

}
func main() {
	cal := make(chan float64)
	go calPmt1(35000, 9.12, 4, cal)
	go calPmt1(35000, 9.12, 10, cal)
	// ans := <-cal
	var input string
	fmt.Scanln(&input)
	fmt.Printf("%.2f", <-cal)
}
