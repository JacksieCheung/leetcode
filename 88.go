package main

// 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，
// 使 nums1 成为一个有序数组。
func merge(nums1 []int, m int, nums2 []int, n int) {
	s1 := nums1[0:m]
	s1 = append(s1, nums2...)
	for i := 0; i < n+m; i++ {
		for v := 0; v < n+m; v++ {
			if s1[i] < s1[v] {
				s1[i], s1[v] = s1[v], s1[i]
			}
		}
	}
}
