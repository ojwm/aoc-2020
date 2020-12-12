package navigation

import (
	"fmt"
	"strings"
)

// Traverse the slope and return the number of trees encountered
func Traverse(slope []string, right int, down int, debug bool) int {
	slopeLength := len(slope) - 1
	slopeWidth := len(slope[0]) - 1
	var slopeMatrix [][]string
	for i, val := range slope {
		slopeMatrix = append(slopeMatrix, strings.Split(val, ""))
		if debug {
			fmt.Printf("%3d", i)
			fmt.Println(" ", slopeMatrix[i])
		}
	}
	x := 0
	y := 0
	trees := 0
	for y <= slopeLength {
		if slopeMatrix[y][x] == "#" {
			trees++
		}
		x = x + right
		if x > slopeWidth {
			x = -1 + x - slopeWidth
		}
		y = y + down
	}
	return trees
}
