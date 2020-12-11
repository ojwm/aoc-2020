package passwordutil

import (
	"strconv"
	"strings"
)

// Validate passwords and return the count of valid entries
func Validate(passwords []string, policy string) int {
	validCount := 0
	for _, val := range passwords {
		entry := strings.Split(val, " ")
		indexes := strings.Split(entry[0], "-")
		index1, _ := strconv.Atoi(indexes[0])
		index2, _ := strconv.Atoi(indexes[1])
		character := entry[1][0:1]
		passwordCharacters := strings.Split(entry[2], "")
		if policy == "sled" {
			count := 0
			for _, val2 := range passwordCharacters {
				if val2 == character {
					count++
				}
			}
			if count >= index1 && count <= index2 {
				validCount++
			}
		} else if policy == "toboggan" {
			if (passwordCharacters[index1-1] == character) != (passwordCharacters[index2-1] == character) {
				validCount++
			}
		}
	}
	return validCount
}
