package leetcode

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	var res []int
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}
	return res
}

func intersectV2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	res := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			res = append(res, nums1[index1])
			index1++
			index2++
		}
	}
	return res
}
