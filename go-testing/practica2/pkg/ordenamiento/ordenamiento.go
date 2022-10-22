package ordenamiento

import (
	"errors"
	"sort"
)

func OrderSlice(nums ...int) ([]int, error) {
	if len(nums) == 0 {
		return nums, errors.New("Debe definin minimo 2 valores para ordenar")
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums, nil
}
