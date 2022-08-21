package main

import "fmt"

func main() {
	myBill := newBill("Mario's bill")

	myBill.addItem("One soup", 4.50)
	myBill.addItem("veg pie", 8.95)
	myBill.addItem("Toffee pudding", 4.95)
	myBill.addItem("coffee", 6.25)
	myBill.updateTip(10)

	fmt.Println(myBill.format())
}
