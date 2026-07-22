package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{1, 2, 3}))
}

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，
// 从最高位到最低位排列。这个大整数不包含任何前导 0。
func plusOne(digits []int) []int {
	// 从最后一位开始遍历
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// 如果遍历完了还走到这一步，说明每一位都是9，新建一个数组，第一位补1
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}
