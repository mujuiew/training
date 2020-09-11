package main

import (
	"fmt"
	"math"
)

var pmt float64

func calPmt1(dis float64, i float64, num int) {
	i = i / 100 / 12
	n := float64(num)
	pmt = dis / ((1 - (1 / math.Pow((1+i), n))) / i)
	fmt.Printf("Payment Amount with Nember of payment %v = %.2f", num, pmt)
	fmt.Println()
}

func main() {
	go calPmt1(35000, 9.12, 4)
	go calPmt1(35000, 9.12, 10)
	var input string
	fmt.Scanln(&input)

}
