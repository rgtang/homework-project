package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1
	// 从第一个开始遍历
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 按照区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 初始化合并后的区间
	merged := [][]int{intervals[0]}
	// 从第二个区间开始遍历
	for i := 1; i < len(intervals); i++ {
		last := &merged[len(merged)-1]
		current := intervals[i]
		if current[0] <= (*last)[1] {
			// 需要合并，合并到右侧区间较大那个
			if current[1] > (*last)[1] {
				(*last)[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}
