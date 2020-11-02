package main

func intersection(nums1 []int, nums2 []int) []int {
	tmp := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				tmp[nums1[i]] = struct{}{}
			}
		}
	}

	var res []int
	for i := range tmp {
		res = append(res, i)
	}
	return res
}
