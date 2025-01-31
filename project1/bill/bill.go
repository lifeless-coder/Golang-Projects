package bill

func Calculate(prev int, cost int) int {
	sum := prev + cost
	return sum
}
func GetCost(m map[string]int, model string) int {
	return m[model]
}
