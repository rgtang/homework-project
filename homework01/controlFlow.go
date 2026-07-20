package main

import "fmt"

// 只出现一次的数字
func singleNumber(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	for num, count := range counts {
		if count == 1 {
			return num
		}
	}
	return 0
}

func main() {
	singleNumber2()
	isPalindrome2()
}

func singleNumber2() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}

func isPalindrome2() {
	fmt.Println(isPalindrome(121))
}

// 判断一个整数是不是回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverse(x)
}

func reverse(x int) int {

	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev
}
