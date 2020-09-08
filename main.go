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
var elements = map[string]map[string]int{
	"Coke": map[string]int{
		"id":     1,
		"price":  12,
		"amount": 0,
	},
	"Water": map[string]int{
		"id":     2,
		"price":  7,
		"amount": 20,
	},
	"Milk": map[string]int{
		"id":     3,
		"price":  22,
		"amount": 20,
	},
}

func main() {

	fmt.Println("ID\t", "Name\t", "Price\t", "Amount\t")
	fmt.Println("----\t", "----\t", "----\t", "----\t")
	for k := range elements {
		if el, ok := elements[k]; ok {
			fmt.Println(el["id"], "\t", k, "\t", el["price"], "\t", el["amount"])
			if productID == 0 || productID < el["id"] {
				productID = el["id"]
			}
		}
	}
	choose()
}
func choose() {
	fmt.Print("\nหากต้องการเพิ่มสิน add ซื้อสินค้ากด get จบการทำงานกด exit: ")
	fmt.Scanf("%s", &ch)
	if ch == "add" {
		addProduct()
		getProduct()
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
	elements[productName] = map[string]int{
		"id":     productID + 1,
		"price":  productPrice,
		"amount": productAmount,
	}
	fmt.Println("ID\t", "Name\t", "Price\t", "Amount\t")
	fmt.Println("----\t", "----\t", "----\t", "----\t")
	for k := range elements {
		if el, ok := elements[k]; ok {
			fmt.Println(el["id"], "\t", k, "\t", el["price"], "\t", el["amount"])
		}
	}
}
func getProduct() {
	fmt.Print("ใส่เงิน (บาทถ้วน): ")
	fmt.Scanf("%d", &money)
	fmt.Println("จำนวนเงินของคุณคือ ", money, " บาท")
	fmt.Println("\nสินค้าที่คุณสามารถเลือกซื้อได้\n ")
	fmt.Println("ID\t", "Name\t", "Price\t", "Amount\t")
	fmt.Println("----\t", "----\t", "----\t", "----\t")
	for k := range elements {
		if el, ok := elements[k]; ok {
			if el["price"] <= money && el["amount"] > 0 {
				fmt.Println(el["id"], "\t", k, "\t", el["price"], "\t", el["amount"])
				x = true
			} else {
				fmt.Println("\nขออภัย! จำนวนเงินของท่าไม่เพียงพอ")
				x = false
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
				case el["price"] <= money && idProduct == el["id"] && el["amount"] > 0:
					fmt.Println("\nสินค้าที่เลือกคือ ", k, " ราคา ", el["price"], " บาท")
					change = float64(money - el["price"])
					fmt.Println("\nกรุณารับเงินทอน : ", change, " บาท")
					// cal()
				case el["price"] > money && idProduct == el["id"]:
					fmt.Println("\nขออภัย! จำนวนเงินของท่าไม่เพียงพอ")
				case idProduct == el["id"] && el["amount"] < 1:
					fmt.Println("\nขออภัย! สินค้าหมด")
				}
			}
		}
	}
}

func exit() {
}
