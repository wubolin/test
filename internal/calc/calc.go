package calc

import "errors"

var ErrEmptyInput = errors.New("at least one number is required")

// Sum returns the total of all numbers.
func Sum(nums ...int) (int, error) {

	if len(nums) == 0 {
		return 0, ErrEmptyInput
	}

	total := 0
	for _, n := range nums {
		total += n
	}
	return total, nil
}

// Multi returns the product of x and y.
func Multi(x, y int) int {
	return x * y
}
