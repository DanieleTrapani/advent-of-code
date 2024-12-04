package days

import (
	"regexp"
	"strings"

	"github.com/DanieleTrapani/advent-of-code-2024/utils"
)

func Day3() int {
	count := 0

	lines := utils.GetData("days/day3.txt")

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	active := true

	for _, line := range lines {
		expressions := re.FindAllStringSubmatch(line, -1)

		for _, expression := range expressions {
			fullMatch := strings.TrimSpace(expression[0])

			// Check for "do()" or "don't()" commands
			if fullMatch == "do()" {
				active = true
				continue
			} else if fullMatch == "don't()" {
				active = false
				continue
			}

			if active && len(expression) == 3 {
				count += utils.ToInt(expression[1]) * utils.ToInt(expression[2])
			}
		}
	}

	return count
}
