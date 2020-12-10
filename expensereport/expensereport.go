package expensereport

// Find expenses that add up to sum and return their product
func Find(expenses []int, sum int, count int) int {
	for i, val1 := range expenses {
		for j, val2 := range expenses {
			if i != j {
				if count == 2 && val1+val2 == sum {
					return val1 * val2
				}
				if count == 3 {
					for k, val3 := range expenses {
						if val1+val2+val3 == sum && j != k {
							return val1 * val2 * val3
						}
					}
				}
			}
		}
	}
	return 0
}
