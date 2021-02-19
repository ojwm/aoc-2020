package boarding

import (
	"fmt"
	"math"
	"strings"
)

// GetHighestSeat from a list of board passes
func GetHighestSeat(passes []string, rows int, cols int) int {
	seatID := 0
	// for _, pass := range passes {
	// seat := strings.Split(pass, "")
	seat := strings.Split(passes[0], "")
	id := getID(seat, rows, cols)
	if id > seatID {
		seatID = id
	}
	// }
	return seatID
}

func getID(seat []string, rows int, cols int) int {
	rowMin, rowMax, colMin, colMax := 0, rows-1, 0, cols-1
	fmt.Println(rowMin, rowMax, colMin, colMax)
	fmt.Println(seat)
	for i, val := range seat {
		midRow := (rowMax - rowMin) / 2
		midCol := (colMax - colMin) / 2
		switch val {
		case "F":
			rowMax = rowMin + getNum(midRow, val)
		case "B":
			rowMin = rowMin + getNum(midRow, val)
		case "L":
			colMax = colMin + getNum(midCol, val)
		case "R":
			colMin = colMin + getNum(midCol, val)
		default:
			return 0
		}
		fmt.Println(rowMin, rowMax, colMin, colMax)
	}
	return 0
}

func getNum(num int, val string) int {
	newNum := float64(num)
	if val == "F" || val == "R" {
		newNum = math.Floor(newNum)
	} else {
		newNum = math.Ceil(newNum)
	}
	return int(newNum)
}
