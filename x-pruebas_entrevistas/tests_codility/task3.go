package main

import (
	"fmt"
	"math"
)

func main() {
	blocks := []int{1, 1}
	fmt.Println(Solution(blocks))
}

func Solution(blocks []int) int {
	// write your code in Go 1.4
	var currentHeight int = blocks[0]
	var max float64
	var i int
	goingUp := false

	// var currentBlockStart, prevBlock int

	if len(blocks) == 2 {
		return len(blocks)
	}

	for i < len(blocks)-1 {
		if (!goingUp && currentHeight >= blocks[i+1]) || (goingUp && currentHeight <= blocks[i+1]) {
			/*if goingUp && currentHeight < blocks[i+1] {
				currentBlockStart = i + 1
			}*/
			currentHeight = blocks[i+1]
			i++
		} else if !goingUp {
			goingUp = true
		} else if goingUp {
			var distance int = i //- prevBlock
			max = math.Max(max, float64(distance))

			/*if i != currentBlockStart {
				prevBlock = currentBlockStart
			} else {
				prevBlock = i
			}*/

			goingUp = false
		}
	}

	var d = i //- prevBlock
	max = math.Max(max, float64(d))

	return int(max)
}
