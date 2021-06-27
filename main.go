package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() string {
	HighAndLow("10 20 42 73")
}

func HighAndLow(in string) string {
	numStrings := strings.Fields(in)
	var nums = []int{}

	for _, i := range numStrings {
		j, _ := strconv.Atoi(i)
		nums = append(nums, j)
	}
	sort.Ints(nums)
	return fmt.Sprintf("%d %d", nums[len(nums)-1], nums[0])
}
