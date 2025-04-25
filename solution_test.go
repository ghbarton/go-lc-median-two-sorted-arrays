package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name string
		data data
		exp  float64
	}{
		{
			name: "returns true",
			data: data{[]int{1, 3}, []int{2}},
			exp:  2.000,
		},
		{
			name: "returns false",
			data: data{[]int{1, 2}, []int{3, 4}},
			exp:  2.500,
		},
		{
			name: "returns false",
			data: data{[]int{1, 2, 3}, []int{4, 5}},
			exp:  3.0,
		},
		{
			name: "returns false",
			data: data{[]int{1, 2, 4, 5}, []int{3}},
			exp:  3.000,
		},
		{
			name: "returns when one array empty",
			data: data{[]int{1, 2, 3}, []int{}},
			exp:  2.000,
		},
		{
			name: "returns when one array empty",
			data: data{[]int{}, []int{1, 2, 3, 4}},
			exp:  2.500,
		},
		{
			name: "length one array effects median of other",
			data: data{[]int{3}, []int{1, 2, 8, 8}},
			exp:  3.000,
		},
		{
			name: "",
			data: data{[]int{1, 2, 3, 4}, []int{3, 4, 5, 6, 7}},
			exp:  4.000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := solution(tc.data)
			assert.Equal(t, tc.exp, audit)
		})
	}

	binarySearchTestCases := []struct {
		name string
		a    []int
		x    int
		exp  int
		err  error
	}{
		{
			name: "returns err",
			a:    []int{0, 2, 3, 4, 5},
			x:    1,
			exp:  0,
			err:  fmt.Errorf("%d was not found in %d", 1, []int{0, 2, 3, 4, 5}),
		},
		{
			name: "finds start",
			a:    []int{1, 2, 3, 4, 5},
			x:    1,
			exp:  0,
			err:  nil,
		},
	}

	for _, tc := range binarySearchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			audit, auditErr := binarySearch(tc.a, tc.x)
			assert.Equal(t, tc.exp, audit)
			assert.Equal(t, tc.err, auditErr)
		})
	}
}
