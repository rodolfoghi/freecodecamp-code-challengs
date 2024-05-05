package main

import (
	"errors"
	"sort"
)

func secondLargest(a []int) (int, error) {
	sort.Ints(a)
	length := len(a)
	if length < 2 {
		return 0, errors.New("array length should be greater than 2")
	}

	return a[length-2], nil
}
