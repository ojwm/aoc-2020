package document

import (
	"regexp"
	"strconv"
	"strings"
)

// CleanDocuments from a list
func CleanDocuments(documents []string) []string {
	var cleanDocuments []string
	var document string
	for _, doc := range documents {
		if len(doc) > 0 {
			document += doc + " "
		} else {
			cleanDocuments = append(cleanDocuments, strings.Trim(document, " "))
			document = ""
		}
	}
	return cleanDocuments
}

// ValidateFormat of documents in a list
func ValidateFormat(documents []string) []string {
	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var validDocuments []string
	for _, doc := range documents {
		valid := true
		for _, field := range mandatory {
			if !strings.Contains(doc, field+":") {
				valid = false
				break
			}
		}
		if valid {
			validDocuments = append(validDocuments, doc)
		}
	}
	return validDocuments
}

// ValidateContent of documents in a list
func ValidateContent(documents []string) []string {
	var validDocuments []string
	for _, doc := range documents {
		if validate(doc) {
			validDocuments = append(validDocuments, doc)
		}
	}
	return validDocuments
}

// validate a document
func validate(document string) bool {
	valid := true
	pairs := strings.Split(document, " ")
	for _, pair := range pairs {
		keyVal := strings.Split(pair, ":")
		switch keyVal[0] {
		case "byr":
			valid = validateInt(keyVal[1], 1920, 2002)
		case "iyr":
			valid = validateInt(keyVal[1], 2010, 2020)
		case "eyr":
			valid = validateInt(keyVal[1], 2020, 2030)
		case "hgt":
			valid = validateHeight(keyVal[1])
		case "hcl":
			valid = validateHair(keyVal[1])
		case "ecl":
			valid = validateEye(keyVal[1])
		case "pid":
			valid = validateID(keyVal[1])
		}
		if !valid {
			break
		}
	}
	return valid
}

func validateHeight(height string) bool {
	if !(len(height) > 3) {
		return false
	}
	val := substr(height, 0, len(height)-2)
	unit := substr(height, len(height)-2, 2)
	var min, max int
	switch unit {
	case "cm":
		min, max = 150, 193
	case "in":
		min, max = 59, 76
	default:
		return false
	}
	return validateInt(val, min, max)
}

func validateHair(colour string) bool {
	match, _ := regexp.MatchString("#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})", colour)
	return match
}

func validateEye(colour string) bool {
	colours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, val := range colours {
		if val == colour {
			return true
		}
	}
	return false
}

func validateID(number string) bool {
	return len(number) == 9 && validateInt(number, 0, 999999999)
}

func validateInt(val string, min int, max int) bool {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return intVal >= min && intVal <= max
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)
	if start >= len(asRunes) {
		return ""
	}
	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}
	return string(asRunes[start : start+length])
}
