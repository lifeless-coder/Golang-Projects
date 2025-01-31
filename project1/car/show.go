package car

import "fmt"

func ShowAvailableCar(m map[string]int) {
	for k, v := range m {
		fmt.Println("Model: ", k, "Price: $", v)
	}
}
