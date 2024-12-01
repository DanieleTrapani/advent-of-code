package days

import (
	"sort"
	"strconv"
	"strings"

	"github.com/DanieleTrapani/advent-of-code-2024/utils"
)

func Day1() int {
	count := 0

	lines := utils.GetData("days/day1.txt")

	// Split into left and right list (convert to int?)
	left_list := make([]int, 0, 1000)
	right_list := make([]int, 0, 1000)

	for _, line := range lines {
		split := strings.Split(line, "   ")
		l, _ := strconv.Atoi(split[0])
		left_list = append(left_list, l)
		r, _ := strconv.Atoi(split[1])
		right_list = append(right_list, r)
	}

	// distance := getDistance(left_list, right_list)

	scores := make(map[int]int, 300)

	for _, val := range left_list {
		if v, ok := scores[val]; ok {
			count += val * v
		} else {
			score := countOccurrences(right_list, val)
			count += val * score
			scores[val] = score
		}
	}

	return count
}

func getDistance(left_list []int, right_list []int) int {
	sum := 0
	// Sort lists
	sort.Slice(left_list, func(i, j int) bool {
		return left_list[i] < left_list[j]
	})
	sort.Slice(right_list, func(i, j int) bool {
		return right_list[i] < right_list[j]
	})

	// Iterate and get distance
	for i := 0; i < len(left_list); i++ {
		distance := utils.Abs(left_list[i] - right_list[i])
		sum += distance
	}
	return sum
}

func countOccurrences(s []int, val int) int {
	sum := 0
	for _, x := range s {
		if x == val {
			sum++
		}
	}
	return sum
}
