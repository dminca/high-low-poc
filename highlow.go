package highlow

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// HighAndLow takes a string of space separated numbers
// and returns the highest and lowest number
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
