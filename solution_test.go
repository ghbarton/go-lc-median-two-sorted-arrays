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
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := solution(tc.data)
			assert.Equal(t, tc.exp, audit)
		})
	}
}
