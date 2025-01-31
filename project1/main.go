package main

import (
	"fmt"
	"project1/bill"
	"project1/order"
)

func main() {

	total := 0
	fmt.Println("WELCOME TO CARSHOP")

	for {
		cost := order.TakeOrder()
		total = bill.Calculate(total, cost)
		fmt.Println("total Cost:", total)
	}
}
