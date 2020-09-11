package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func main2() {
	pList := [][]string{
		[]string{"1", "Coke", "-", "-"},
		[]string{"2", "Water", "-", "-"},
		[]string{"3", "Milk", "-", "-"},
	}
	// The players take turns.
	pList[0][3] = "1"
	pList[0][2] = "12"

	pList[1][3] = "O"
	pList[1][2] = "7"

	pList[2][3] = "2"
	pList[2][2] = "22"

	m := map[string][]string{
		"Coke":  pList[0],
		"Water": pList[1],
		"Milk":  pList[2],
	}
	fmt.Println(m)
	for i := 0; i < len(pList); i++ {
		fmt.Println(pList[i][0], "\t")
	}
	fmt.Println("----------------------------------------------------------------------------------")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, " %s\t%s\t%s\t%s\t", "ID", "Name", "Price", "Amount")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "----", "----", "----", "----")

	for i := 0; i < len(pList); i++ {
		fmt.Fprintf(w, "\n%s\n", strings.Join(pList[i], "\t"))
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
Loop:
	for i := 0; i < len(pList); i++ {
		Price, _ := strconv.Atoi(pList[i][2])
		Amount, _ := strconv.Atoi(pList[i][3])
		switch {
		case (Price <= money) && (Amount > 0):
			fmt.Fprintf(w, "\n%s\n", strings.Join(pList[i], "\t"))
			x = true
		case money < Price:
			fmt.Println("ขออภัย! จำนวนเงินของคุณไม่เพิ่งพอในการเลือกซื้อสินค้า")
			fmt.Println("เงินทอน ", money, " บาท")
			break Loop
		}
	}
	w.Flush()
	if x {
		fmt.Print("\n-----\nกรุณาใส่รหัสสินค้า : ")
		fmt.Scanf("%d", &idProduct)
		fmt.Print("รหัสสินค้า : ", idProduct)
		for i := 0; i < len(pList); i++ {
			ID, _ := strconv.Atoi(pList[i][0])
			Price, _ := strconv.Atoi(pList[i][2])
			Amount, _ := strconv.Atoi(pList[i][3])
			switch {
			case Price <= money && idProduct == ID && Amount > 0:
				fmt.Println("\nสินค้าที่เลือกคือ ", pList[i][1], " ราคา ", Price, " บาท")
				change = float64(money - Price)
				fmt.Println("\nกรุณารับเงินทอน : ", change, " บาท")
				// cal()
				// break InputIdProduct
			case Price > money && idProduct == ID:
				fmt.Println("\nขออภัย! จำนวนเงินของท่าไม่เพียงพอ")
				// break InputIdProduct
			case idProduct == ID && Amount < 1:
				fmt.Println("\nขออภัย! สินค้าหมด")
				// break InputIdProduct
			}
		}
	}

}
