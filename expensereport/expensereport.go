package expensereport

// Find two expenses that add up to sum and return their product
func Find(expenses []int, sum int) int {
	for _, value := range expenses {
		i := 0
		for range expenses {
			if value+expenses[i] == sum {
				return value * expenses[i]
			}
			i++
		}
	}
	return 0
}
