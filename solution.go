package main

import (
	"fmt"
	"math"
)

type data struct {
	nums1 []int
	nums2 []int
}

func solution(data data) float64 {
	a := []int{} // always larger
	b := []int{}
	if len(data.nums2) == 0 {
		return getMedianOfArray(data.nums1)
	}
	if len(data.nums1) == 0 {
		return getMedianOfArray(data.nums2)
	}
	if data.nums1[0] > data.nums2[0] {
		a = data.nums2
		b = data.nums1
	} else {
		a = data.nums1
		b = data.nums2
	}
	if a[len(a)-1] < b[0] { // consecutive arrays
		mid := float64(len(a)+len(b)) / 2.0
		point := getPointFromConsecutiveArrays(a, b, int(math.Ceil(float64(mid))))
		if mid != math.Floor(mid) {
			return point
		}
		point2 := getPointFromConsecutiveArrays(a, b, int(math.Round(float64(mid)+1)))
		return (point + point2) / 2.0
	}
	if b[len(b)-1] < a[len(a)-1] {
		mid := int(math.Ceil(float64(len(b)) / 2.0))
		return float64(b[mid-1])
	}
	mid := float64(len(a)+len(b)) / 2.0 // assume single point median for now
	if int(mid) < len(a) {
		midVal := a[int(mid)]
		if midVal < b[0] {
			return float64(midVal)
		}
	} else {
		midVal := b[int(mid)-len(a)]
		if a[len(a)-1] < midVal {
			return float64(midVal)
		}
	}

	//if a[len(a)-1] <= b[len(b)-1] && a[0] <= a[0] {
	//	am := getMedianOfArray(a)
	//	bm := getMedianOfArray(b)
	//
	//}
	return -1.0
}

func binarySearch(a []int, x int) (int, error) {
	if len(a) == 0 {
		return -1, fmt.Errorf("cannot search empty array")
	}
	if len(a) == 1 {
		if a[0] == x {
			return 0, nil
		}
		return -1, fmt.Errorf("%d was not found in %d", x, a)
	}
	r := len(a) - 1
	m := r / 2
	l := 0
	for l < r {
		if x < a[m] {
			r = m
		} else {
			l = m
		}
		if a[l] == x {
			return l, nil
		}
		if a[r] == x {
			return r, nil
		}
		m = l + (r-l)/2
	}
	return -1, fmt.Errorf("%d was not found in %d", x, a)
}
func getPointFromConsecutiveArrays(a []int, b []int, p int) float64 {
	p = p - 1
	if len(a) <= p {
		return float64(b[p-len(a)])
	}
	return float64(a[p])
}

func getMedianOfArray(a []int) float64 {
	mid := float64(len(a)) / 2.0
	midFloor := math.Floor(mid)
	point := a[int(midFloor)]
	if mid != midFloor {
		return float64(point)
	}
	point2 := a[int(math.Floor(midFloor-1))]
	return float64(point+point2) / 2.0
}
