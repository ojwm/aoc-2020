package passwordutil

import (
	"strconv"
	"strings"
)

// Validate passwords and return the count of valid entries
func Validate(passwords []string) int {
	validCount := 0
	for _, val := range passwords {
		entry := strings.Split(val, " ")
		countRange := strings.Split(entry[0], "-")
		minCount, _ := strconv.Atoi(countRange[0])
		maxCount, _ := strconv.Atoi(countRange[1])
		character := entry[1][0:1]
		passwordCharacters := strings.Split(entry[2], "")
		count := 0
		for _, val2 := range passwordCharacters {
			if val2 == character {
				count++
			}
		}
		if count >= minCount && count <= maxCount {
			validCount++
		}
	}
	return validCount
}
