package main

import "sort"

func main() {
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	// 有序数组才能二分查找
	sort.Ints(potions)
	ans := make([]int, len(spells))

	for i, spell := range spells {
		// sort.SearchInts 返回第一个大于等于 target 的索引，就说明该索引后续的元素都满足条件
		// (int(success)-1)/spell+1 表示满足 success 的最小药水能量值
		// 用 len(potions) 减去这个索引，就是满足条件的药水数量
		ans[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/spell+1)
	}

	return ans
}
