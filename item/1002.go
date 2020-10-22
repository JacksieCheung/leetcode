package main

// 查找常用字符
// 给ｎ个字符串的数组，找共同的字符。
// 每次遍历一个字符串，然后用新数组存储，再对比原数组判断大小，保留小的
// 走完之后每次取最小，最后就是一共出现的次数。
// 考察int 类型索引和字符变量之间的转换，通过'a'+i表示二十六个字母。
func commonChars(A []string) []string {
	nums := [26]int{}
	for i := 0; i < len(nums); i++ {
		nums[i] = math.MaxInt64
	}
	for _, word := range A {
		num := [26]int{}
		for _, c := range word {
			num[c-'a']++
		}
		for i, v := range num {
			if nums[i] > v {
				nums[i] = v
			}
		}
	}
	// fmt.Println(nums)
	var res []string
	for i, v := range nums {
		for k := 0; k < v; k++ {
			res = append(res, string(byte(i)+'a'))
		}
	}
	return res
}
