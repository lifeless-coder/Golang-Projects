package order

import (
	"fmt"
	"project1/bill"
	"project1/car"
)

func inputModel() string {
	fmt.Println("Input your choice model")
	var model string
	fmt.Scanln(&model)
	return model
}
func TakeOrder() int {

	fmt.Println("Select car brand")
	fmt.Println("1.Toyota")
	fmt.Println("2.Honda")
	fmt.Println("3.BMW")
	fmt.Println("4.Audi")
	fmt.Println("Enter your choice")

	var choice int

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		car.ShowAvailableCar(car.Toyota)
		model := inputModel()
		cost := bill.GetCost(car.Toyota, model)
		return cost
	case 2:
		car.ShowAvailableCar(car.Honda)
		model := inputModel()
		cost := bill.GetCost(car.Honda, model)
		return cost
	case 3:
		car.ShowAvailableCar(car.BMW)
		model := inputModel()
		cost := bill.GetCost(car.BMW, model)
		return cost
	case 4:
		car.ShowAvailableCar(car.Audi)
		model := inputModel()
		cost := bill.GetCost(car.Audi, model)
		return cost
	default:
		fmt.Println("Invalid choice")
		return 0
	}
	return 0
}
