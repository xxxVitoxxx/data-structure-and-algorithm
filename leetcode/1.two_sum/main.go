package main

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		idx, ok := m[target-v]
		if ok {
			return []int{idx, i}
		}

		m[v] = i
	}
	return nil
}

func main() {
	tests := map[string]struct {
		input  []int
		target int
		output []int
	}{
		"example 1": {
			input:  []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		"example 2": {
			input:  []int{3, 1, 2, 4},
			target: 6,
			output: []int{2, 3},
		},
		"example 3": {
			input:  []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}

	for k, v := range tests {
		fmt.Printf("%s -> ", k)
		result := twoSum(v.input, v.target)
		fmt.Printf("%v \n", reflect.DeepEqual(result, v.output))
	}
}
