package main

// 全排列有重复
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	dfs(nums, []int{}, &res)
	fmt.Println(res)
	return res
}

func dfs(numsLeft []int, track []int, res *[][]int) {
	if len(numsLeft) == 0 { //判断是否停下来
		*res = append(*res, track)
		return
	} else {
		for i := 0; i < len(numsLeft); i++ {
			if i != 0 && numsLeft[i] == numsLeft[i-1] {
				continue
			}
			tmp := append([]int{}, numsLeft...)
			trackcopy := append([]int{}, track...)
			dfs(append(tmp[:i], tmp[i+1:]...), append(trackcopy, numsLeft[i]), res) //记录走过的路和选择道路
		}
	}
}
