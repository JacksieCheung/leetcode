package main

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
func plusOne(digits []int) []int {
	l := len(digits)
	digits[l-1]++
	if digits[l-1] != 10 {
		return digits
	}
	if l == 1 {
		digits[0] = 0
		digits = append(digits, 1)
		digits[0], digits[1] = digits[1], digits[0]
		return digits
	}
	//要进位的情况
	for i := l - 1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i-1]++
		}

		if digits[0] == 10 {
			//fmt.Println("append")
			digits = append(digits, 1)
			break
		}
	}
	//fmt.Println(digits)
	//把10全部换为0
	l2 := len(digits)
	for i := 0; i < l2; i++ {
		if digits[i] == 10 {
			digits[i] = 0
		}
	}
	//把1换到前面来进位
	if digits[0] == 0 {
		digits[0], digits[l2-1] = digits[l2-1], digits[0]
	}
	return digits
}
