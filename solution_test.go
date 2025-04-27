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
	slotSearchTestCases := []struct {
		name string
		a    []int
		x    int
		exp  int
		err  error
	}{
		{
			name: "empty array error",
			a:    []int{},
			x:    1,
			exp:  -1,
			err:  fmt.Errorf("cannot search empty array"),
		},
		{
			name: "single item array ",
			a:    []int{5},
			x:    1,
			exp:  0,
			err:  nil,
		},
		{
			name: "returns 1",
			a:    []int{0, 2, 3, 4, 5},
			x:    1,
			exp:  1,
			err:  nil,
		},
		{
			name: "returns 1",
			a:    []int{0, 1, 3, 4, 5},
			x:    1,
			exp:  1,
			err:  nil,
		},
		{
			name: "returns 4",
			a:    []int{0, 1, 3, 4, 5},
			x:    5,
			exp:  4,
			err:  nil,
		},
		{
			name: "returns 5",
			a:    []int{0, 1, 3, 4, 5},
			x:    6,
			exp:  5,
			err:  nil,
		},
	}

	for _, tc := range slotSearchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			audit, auditErr := slotSearch(tc.a, tc.x)
			assert.Equal(t, tc.exp, audit)
			assert.Equal(t, tc.err, auditErr)
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
			name: "empty array error",
			a:    []int{},
			x:    1,
			exp:  -1,
			err:  fmt.Errorf("cannot search empty array"),
		},
		{
			name: "returns err",
			a:    []int{0, 2, 3, 4, 5},
			x:    1,
			exp:  -1,
			err:  fmt.Errorf("%d was not found in %d", 1, []int{0, 2, 3, 4, 5}),
		},
		{
			name: "returns err",
			a:    []int{0},
			x:    1,
			exp:  -1,
			err:  fmt.Errorf("%d was not found in %d", 1, []int{0}),
		},
		{
			name: "finds start",
			a:    []int{1, 2, 3, 4, 5},
			x:    1,
			exp:  0,
			err:  nil,
		},
		{
			name: "finds end",
			a:    []int{1, 2, 3, 4, 5},
			x:    5,
			exp:  4,
			err:  nil,
		},
		{
			name: "finds mid",
			a:    []int{1, 2, 3, 4, 5},
			x:    3,
			exp:  2,
			err:  nil,
		},
		{
			name: "finds mid",
			a:    []int{0, 0, 3, 4, 5},
			x:    3,
			exp:  2,
			err:  nil,
		},
		{
			name: "finds mid",
			a:    []int{0, 0, 3, 4, 5},
			x:    0,
			exp:  0,
			err:  nil,
		},
		{
			name: "single element array - match",
			a:    []int{5},
			x:    5,
			exp:  0,
			err:  nil,
		},
		{
			name: "large sorted array - find first",
			a:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			x:    1,
			exp:  0,
			err:  nil,
		},
		{
			name: "large sorted array - find last",
			a:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			x:    10,
			exp:  9,
			err:  nil,
		},
		{
			name: "duplicates - target at beginning",
			a:    []int{2, 2, 2, 3, 4},
			x:    2,
			exp:  2,
			err:  nil,
		},
		{
			name: "all elements same, target present",
			a:    []int{7, 7, 7, 7, 7},
			x:    7,
			exp:  2,
			err:  nil,
		},
		{
			name: "all elements same, target missing",
			a:    []int{7, 7, 7, 7, 7},
			x:    8,
			exp:  -1,
			err:  fmt.Errorf("%d was not found in %d", 8, []int{7, 7, 7, 7, 7}),
		},
		{
			name: "negative numbers - target present",
			a:    []int{-10, -5, 0, 5, 10},
			x:    -5,
			exp:  1,
			err:  nil,
		},
		{
			name: "negative numbers - target missing",
			a:    []int{-10, -5, 0, 5, 10},
			x:    -6,
			exp:  -1,
			err:  fmt.Errorf("%d was not found in %d", -6, []int{-10, -5, 0, 5, 10}),
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
