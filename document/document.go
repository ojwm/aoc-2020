package document

import (
	"strings"
)

// validate a document
// func validate (document )

func cleanDocuments(documents []string) []string {
	var cleanDocuments []string
	var document string
	for _, val := range documents {
		if len(val) > 0 {
			document += val + " "
		} else {
			cleanDocuments = append(cleanDocuments, strings.Trim(document, " "))
			document = ""
		}
	}
	return cleanDocuments
}

// Validate a list of documents
func Validate(documents []string) int {
	cleanedDocuments := cleanDocuments(documents)
	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := 0
	for _, val := range cleanedDocuments {
		flag := true
		for _, val2 := range mandatory {
			if !strings.Contains(val, val2+":") {
				flag = false
			}
		}
		if flag {
			valid++
		}
	}
	return valid
}
