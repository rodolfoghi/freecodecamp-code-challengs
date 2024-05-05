package main

import "testing"

func TestSecondLargest_Tests(t *testing.T) {
	t.Parallel() // marks TLog as capable of running in parallel with other tests

	var tests = []struct {
		expected int
		values   []int
	}{
		{expected: 40, values: []int{10, 40, 30, 20, 50}},
		{expected: 105, values: []int{25, 143, 89, 13, 105}},
		{expected: 23, values: []int{54, 23, 11, 17, 10}},
	}

	for _, test := range tests {
		actual, err := secondLargest(test.values)
		if err != nil {
			t.Errorf(err.Error())
		}
		if actual != test.expected {
			t.Errorf("got %d, want %d", actual, test.expected)
		}
	}
}
