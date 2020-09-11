package main

import (
	"fmt"
)

var money int
var ch string
var x bool
var idProduct int
var change float64
var productName string
var productID int
var productPrice int
var productAmount int

type product struct {
	//ID     int
	Name   string
	Price  int
	Amount int
}

var elements = map[int]product{
	1: {"Coke", 18, 10},
	2: {"Water", 7, 10},
	5: {"Milk", 12, 10},
}

func main() {

	fmt.Println("ID\t", "Name\t", "Price\t", "Amount\t")
	fmt.Println("----\t", "----\t", "----\t", "----\t")
	v := make([]product, 0, len(elements))
	for _, value := range elements {
		v = append(v, value)
	}
	for k := range elements {

		if el, ok := elements[k]; ok {
			fmt.Println(k, "\t", el.Name, "\t", el.Price, "\t", el.Amount)

		}
	}
	choose()
}
func choose() {
	for k := range elements {
		if _, ok := elements[k]; ok {
			if productID == 0 || productID < k {
				productID = k
			}
		}
	}
	fmt.Print("\nหากต้องการเพิ่มสิน add ซื้อสินค้ากด get จบการทำงานกด exit: ")
	fmt.Scanf("%s", &ch)
	if ch == "add" {
		addProduct()
		choose()
	} else if ch == "get" {
		getProduct()
		choose()
	} else if ch == "exit" {
		exit()
	} else {
		choose()
	}
}
func addProduct() {
	fmt.Print("ชื่อสินค้า: ")
	fmt.Scanf("%s", &productName)
	fmt.Print("ราคาสินค้า: ")
	fmt.Scanf("%d", &productPrice)
	fmt.Print("จำนวนสินค้า: ")
	fmt.Scanf("%d", &productAmount)
	newID := productID + 1
	elements[newID] = product{
		productName, productPrice, productAmount,
	}

	fmt.Println("ID\t", "Name\t", "Price\t", "Amount\t")
	fmt.Println("----\t", "----\t", "----\t", "----\t")
	for k := range elements {
		if el, ok := elements[k]; ok {
			fmt.Println(k, "\t", el.Name, "\t", el.Price, "\t", el.Amount)
		}
	}
}
func getProduct() {
	fmt.Print("\nใส่เงิน (บาทถ้วน): ")
	fmt.Scanf("%d", &money)
	fmt.Println("จำนวนเงินของคุณคือ ", money, " บาท")
	fmt.Println("\nสินค้าที่คุณสามารถเลือกซื้อได้\n ")
	fmt.Println("ID\t", "Name\t", "Price\t", "Amount\t")
	fmt.Println("----\t", "----\t", "----\t", "----\t")
	x = false
	for k := range elements {
		if el, ok := elements[k]; ok {
			if el.Price <= money && el.Amount > 0 {
				fmt.Println(k, "\t", el.Name, "\t", el.Price, "\t", el.Amount)
				x = true
			}
		}
	}
	fmt.Print("\n")
	if x {
		fmt.Print("\n-----\nกรุณาใส่รหัสสินค้า : ")
		fmt.Scanf("%d", &idProduct)
		fmt.Print("รหัสสินค้า : ", idProduct)

		for k := range elements {

			if el, ok := elements[k]; ok {
				switch {
				case el.Price <= money && idProduct == k && el.Amount > 0:
					fmt.Println("\nสินค้าที่เลือกคือ ", k, " ราคา ", el.Price, " บาท")
					var p = *&el.Amount - 1
					el.Amount = p
					// el.Amount--
					elements[k] = el
					fmt.Println(p)

					change = float64(money - el.Price)
					fmt.Println("\nกรุณารับเงินทอน : ", change, " บาท")
					// cal()
				case el.Price > money && idProduct == k:
					fmt.Println("\nขออภัย! จำนวนเงินของท่าไม่เพียงพอ")
				case idProduct == k && el.Amount < 1:
					fmt.Println("\nขออภัย! สินค้าหมด")
				}
			}
		}
	} else {
		fmt.Println("\nขออภัย! จำนวนเงินของท่าไม่เพียงพอ")
	}
}

func exit() {
}
