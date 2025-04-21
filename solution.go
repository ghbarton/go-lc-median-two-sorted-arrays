package main

import "math"

type data struct {
	nums1 []int
	nums2 []int
}

func solution(data data) float64 {
	a := []int{} // always larger
	b := []int{}
	if len(data.nums2) == 0 { // todo implement for no array
		return float64(data.nums2[mid])
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
		if mid/math.Floor(mid) != 1 {
			return point
		}
		point2 := getPointFromConsecutiveArrays(a, b, int(math.Round(float64(mid)+1)))
		return (point + point2) / 2.0
	}
	if b[len(b)-1] < a[len(a)-1] {
		mid := int(math.Ceil(float64(len(b)) / 2.0))
		return float64(b[mid-1])
	}
	return 0.0
}

func getPointFromConsecutiveArrays(a []int, b []int, p int) float64 {
	p = p - 1
	if len(a) <= p {
		return float64(b[p-len(a)])
	}
	return float64(a[p])
}

func getMedianOfArray(a []int) float64 {
	mid := len(a) / 2.0
	point := math.Ceil(mid)
	if mid/math.Floor(mid) != 1 {
		return point
	}
	point2 := getPointFromConsecutiveArrays(a, b, int(math.Round(float64(mid)+1)))
	return (point + point2) / 2.0
}
