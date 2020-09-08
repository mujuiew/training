package main

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

//Product type
type Product struct {
	Flag   bool
	ID     int
	Name   string
	Price  int
	Amount int
}

func checknull() {
	test := Product{
		Flag:   false,
		ID:     0,
		Name:   "",
		Price:  0,
		Amount: 0,
	}
	fmt.Println("test")
	if test.Flag != false {
		fmt.Println("true :", test)
	} else {
		fmt.Println("false :", test)
	}
}
func main1() {
	checknull()
	pList := [3]Product{}
	pList[0] = Product{
		ID:     1,
		Name:   "Coca-Cola",
		Price:  20,
		Amount: 0,
	}
	pList[1] = Product{
		ID:     2,
		Name:   "Water",
		Price:  7,
		Amount: 20,
	}
	pList[2] = Product{
		ID:     3,
		Name:   "Milk",
		Price:  18,
		Amount: 20,
	}
	fmt.Println("Product All")
	// Product All
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, " %s\t%s\t%s\t%s\t", "ID", "Name", "Price", "Amount")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "----", "----", "----", "----")
	for _, i := range pList {
		fmt.Fprintf(w, "\n %d\t%q\t%d\t%d\t", i.ID, i.Name, i.Price, i.Amount)
	}
	w.Flush()
	fmt.Println("\n ")

	fmt.Print("ใส่เงิน (บาทถ้วน): ")
	fmt.Scanf("%d", &money)
	fmt.Println("จำนวนเงินของคุณคือ ", money, " บาท")
	fmt.Println("\nสินค้าที่คุณสามารถเลือกซื้อได้\n-----")

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, " %s\t%s\t%s\t%s\t", "ID", "Name", "Price", "Amount")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "----", "----", "----", "----")
	for _, i := range pList {
		switch {
		case i.Price <= money && i.Amount > 0:
			fmt.Fprintf(w, "\n %d\t%q\t%d\t%d\t", i.ID, i.Name, i.Price, i.Amount)

		case money < i.Price:
			fmt.Println("ขออภัย! จำนวนเงินของคุณไม่เพิ่งพอในการเลือกซื้อสินค้า")
			fmt.Println("เงินทอน ", money, " บาท")
		}
	}
	w.Flush()

	fmt.Print("\n-----\nกรุณาใส่รหัสสินค้า : ")
	fmt.Scanf("%d", &idProduct)
	fmt.Print("รหัสสินค้า : ", idProduct)

	// InputIdProduct:
	for _, i := range pList {

		switch {
		case i.Price <= money && idProduct == i.ID && i.Amount > 0:
			fmt.Println("\nสินค้าที่เลือกคือ ", i.Name, " ราคา ", i.Price, " บาท")
			change = float64(money - i.Price)
			fmt.Println("\nกรุณารับเงินทอน : ", change, " บาท")
			cal()
			// break InputIdProduct
		case i.Price > money && idProduct == i.ID:
			fmt.Println("\nขออภัย! จำนวนเงินของท่าไม่เพียงพอ")
			// break InputIdProduct
		case idProduct == i.ID && i.Amount < 1:
			fmt.Println("\nขออภัย! สินค้าหมด")
			// break InputIdProduct
		}
	}

}

func cal() {
	if change >= 10 {
		sum10 := change / 10
		sum10s := math.Floor(sum10) * 10
		change = change - sum10s
		fmt.Println("มีจำนวนเหรียญ 10 = ", math.Floor(sum10), " เหรียญ")
	}
	if change >= 5 {
		sum5 := change / 5
		sum5s := math.Floor(sum5) * 5
		change = change - sum5s
		fmt.Println("มีจำนวนเหรียญ 5 = ", math.Floor(sum5), " เหรียญ")
	}
	if change >= 2 {
		sum2 := change / 2
		sum2s := math.Floor(sum2) * 2
		change = change - sum2s
		fmt.Println("มีจำนวนเหรียญ 2 = ", math.Floor(sum2), " เหรียญ")
	}
	if change >= 1 {
		sum1 := change / 1
		sum1s := math.Floor(sum1) * 1
		change = change - sum1s
		fmt.Println("มีจำนวนเหรียญ 1 = ", math.Floor(sum1), " เหรียญ")
	}
}
