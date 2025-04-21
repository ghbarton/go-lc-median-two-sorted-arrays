package main

import (
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
			exp:  3.000,
		},
		{
			name: "returns when one array empty",
			data: data{[]int{}, []int{1, 2, 3, 4}},
			exp:  3.000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := solution(tc.data)
			assert.Equal(t, tc.exp, audit)
		})
	}
}
