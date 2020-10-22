package main

// 全排列
func permute(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, []int{}, &res)
	fmt.Println(res)
	return res
}

func dfs(numsLeft []int, track []int, res *[][]int) {
	if len(numsLeft) == 0 {
		*res = append(*res, track)
		return
	} else {
		for i, v := range numsLeft {
			tmp := make([]int, i)
			copy(tmp, numsLeft[:i])
			dfs(append(tmp, numsLeft[i+1:]...), append(track, v), res)
		}
	}
}
