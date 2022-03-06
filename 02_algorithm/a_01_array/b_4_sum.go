package main

/**
Author: xym
Date: 2021/5/23 20:21
Project: golearning
Description:
*/

/*
4. 三数之和 中等
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]
提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))  // [[-1 -1 2] [-1 0 1]]
	fmt.Println(badTwoSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))       // [[-2 0 2]]
}

// 处理结果与元素原顺序无关，可排序预处理，方便去重
// 使用双指针遍历后部分剩余数组
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // [-4 -1 -1 0 1 2] // 预排序有 2 个好处：去重 & 指导双指针的下一步方向
	n := len(nums)
	var res [][]int
	for i, num := range nums {
		if num > 0 {
			break // 优化，再往后三个正数和不可能为 0
		}

		// 第一层遍历数向前去重
		if i > 0 && nums[i] == nums[i-1] { // 因为双指针从 i 之后取，不能使用 nums[i] == nums[i+1] 向后去重
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := num + nums[l] + nums[r]
			switch {
			case sum > 0:
				r--
			case sum < 0:
				l++
			default:
				res = append(res, []int{num, nums[l], nums[r]})
				// 第二层候选数向后去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}

// twoSum 的思路，不好
func badTwoSum(nums []int) [][]int {
	// 避开全是 0 的 case     // ugly
	if len(nums) >= 3 {
		allZero := true
		for _, num := range nums {
			if num != 0 {
				allZero = false
			}
		}
		if allZero {
			return [][]int{{0, 0, 0}}
		}
	}

	n := len(nums)
	num2index := make(map[int]int, n)
	for i, num := range nums {
		num2index[num] = i
	}

	// 获取三元组
	var res [][]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			remain := 0 - (nums[i] + nums[j])
			if k, ok := num2index[remain]; ok && j != k && i != k {
				res = append(res, []int{nums[i], nums[j], remain})
			}
		}
	}

	// 剔除重复的三元组
	m := make(map[string][]int)
	for i := range res {
		sort.Ints(res[i])
		m[intStr(res[i])] = res[i]
	}

	var arrs [][]int
	for _, arr := range m {
		arrs = append(arrs, arr)
	}
	return arrs
}

// trick    // 使得整数数组能做 map 的 key
func intStr(nums []int) string {
	str := ""
	for _, num := range nums {
		str += fmt.Sprintf("%d_", num)
	}
	return str
}
