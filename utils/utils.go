package utils

import "sort"

func GetUnionOfTwoLists(listA_ []int, listB_ []int) []int {
	listA := listA_
	listB := listB_

	union := make(map[int]bool)

	// 将列表 A 中的元素添加到并集
	for _, num := range listA {
		union[num] = true
	}

	// 将列表 B 中的元素添加到并集
	for _, num := range listB {
		union[num] = true
	}

	// 将并集转换回切片
	result := make([]int, 0, len(union))
	for num := range union {
		result = append(result, num)
	}

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	return result
}
