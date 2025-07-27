package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main1() {
	//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	a := []int{1, 2, 2, 3, 3}
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			fmt.Println(k)
		}
	}
	//考察：数字操作、条件判断,题目：判断一个整数是否是回文数
	b := 12321
	c := huiwen(b)
	fmt.Println(c)
	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	s := "()[]{}"
	fmt.Println(isValid(s))

	//查找字符串数组中的最长公共前缀
	arr := []string{"fl", "flo", "float"}
	fmt.Println(longestCommonPrefix(arr))

	//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
	f := []int{1, 2, 9}
	fmt.Println(plusOne(f))

	//给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
	// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
	// 可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，
	// 一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
	nums1 := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums1))

	//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
	nums := []int{9, 7, 2}
	target := 9
	fmt.Println(twoSum(nums, target))

}

func huiwen(b int) bool {
	c := strconv.Itoa(b)
	for i := 0; i < len(c)/2; i++ {
		if c[i] != c[len(c)-1-i] {
			return false
		}
	}
	return true
}

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1] // 弹出栈顶
		}
	}

	return len(stack) == 0
}

func longestCommonPrefix(strs []string) string {
	// 如果字符串数组为空，直接返回
	if len(strs) == 0 {
		return ""
	}
	// 初始化最长公共前缀为第一个字符串
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始处理
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++

		// 如果没有进位，直接返回
		if digits[i] < 10 {
			return digits
		}

		// 有进位，当前位设为0，继续处理前一位
		digits[i] = 0
	}

	// 如果所有位都处理完仍有进位，在最前面插入1
	return append([]int{1}, digits...)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // 值到索引的映射
	for i, num := range nums {
		complement := target - num
		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}
