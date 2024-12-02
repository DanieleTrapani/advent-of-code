package days

import (
	"strings"

	"github.com/DanieleTrapani/advent-of-code-2024/utils"
)

func Day2() int {
	count := 0

	lines := utils.GetData("days/day2test.txt")

Outer:
	for _, line := range lines {
		ints := utils.MapToInt(strings.Split(line, " "))
		ascending := false
		if ints[0] < ints[1] {
			ascending = true
		}

		for i := 0; i < len(ints)-1; i++ {
			prev := ints[i]
			next := ints[i+1]
			if (prev < next) != ascending {
				continue Outer
			}

			diff := utils.Abs(prev - next)
			if diff < 1 || diff > 3 {
				continue Outer
			}
		}
		count++
	}

	return count
}
