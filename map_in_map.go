package main

import (
	"fmt"
)

var elements2 = map[string]map[string]int{
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

func main4() {

	fmt.Println("ID\t", "Name\t", "Price\t", "Amount\t")
	fmt.Println("----\t", "----\t", "----\t", "----\t")
	for k := range elements2 {
		if el, ok := elements2[k]; ok {
			fmt.Println(el["id"], "\t", k, "\t", el["price"], "\t", el["amount"])
			if productID == 0 || productID < el["id"] {
				productID = el["id"]
			}
		}
	}
	chooseMap()
}
func chooseMap() {
	fmt.Print("\nหากต้องการเพิ่มสิน add ซื้อสินค้ากด get จบการทำงานกด exit: ")
	fmt.Scanf("%s", &ch)
	if ch == "add" {
		addProductMap()
		chooseMap()
	} else if ch == "get" {
		getProductMap()
		chooseMap()
	} else if ch == "exit" {
		exitMap()
	} else {
		chooseMap()
	}
}
func addProductMap() {
	fmt.Print("ชื่อสินค้า: ")
	fmt.Scanf("%s", &productName)
	fmt.Print("ราคาสินค้า: ")
	fmt.Scanf("%d", &productPrice)
	fmt.Print("จำนวนสินค้า: ")
	fmt.Scanf("%d", &productAmount)
	elements2[productName] = map[string]int{
		"id":     productID + 1,
		"price":  productPrice,
		"amount": productAmount,
	}
	fmt.Println("ID\t", "Name\t", "Price\t", "Amount\t")
	fmt.Println("----\t", "----\t", "----\t", "----\t")
	for k := range elements2 {
		if el, ok := elements2[k]; ok {
			fmt.Println(el["id"], "\t", k, "\t", el["price"], "\t", el["amount"])
		}
	}
}
func getProductMap() {
	fmt.Print("ใส่เงิน (บาทถ้วน): ")
	fmt.Scanf("%d", &money)
	fmt.Println("จำนวนเงินของคุณคือ ", money, " บาท")
	fmt.Println("\nสินค้าที่คุณสามารถเลือกซื้อได้\n ")
	fmt.Println("ID\t", "Name\t", "Price\t", "Amount\t")
	fmt.Println("----\t", "----\t", "----\t", "----\t")
	x = false
	for k := range elements2 {
		if el, ok := elements2[k]; ok {
			if el["price"] <= money && el["amount"] > 0 {
				fmt.Println(el["id"], "\t", k, "\t", el["price"], "\t", el["amount"])
				x = true
			}
		}
	}
	fmt.Print("\n")
	if x {
		fmt.Print("\n-----\nกรุณาใส่รหัสสินค้า : ")
		fmt.Scanf("%d", &idProduct)
		fmt.Print("รหัสสินค้า : ", idProduct)
		for k := range elements2 {
			if el, ok := elements2[k]; ok {
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
	} else {
		fmt.Println("\nขออภัย! จำนวนเงินของท่าไม่เพียงพอ")
	}
}

func exitMap() {
}
