## 2300

### 题
给你两个正整数数组 spells 和 potions ，长度分别为 n 和 m ，其中 spells[i] 表示第 i 个咒语的能量强度，potions[j] 表示第 j 瓶药水的能量强度。

同时给你一个整数 success 。一个咒语和药水的能量强度 相乘 如果 大于等于 success ，那么它们视为一对 成功 的组合。

请你返回一个长度为 n 的整数数组 pairs，其中 pairs[i] 是能跟第 i 个咒语成功组合的 药水 数目。

### 题解
```go
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
```
这里的二分查找用的轮子，有时候是可以用用的。

以下是自己实现的二分查找函数：
```go
func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions)
    ans := make([]int, len(spells))
    
    for i, spell := range spells {
        // 计算目标值（关键逻辑）
        target := (int(success)-1)/spell + 1
        // 使用自定义二分查找
        index := mySearchInts(potions, target)
        ans[i] = len(potions) - index
    }
	
    return ans
}

func MySearchInts(potions []int, target int) int {
    left, right := 0, len(potions)
    for left < right {
        mid := left + (right - left) / 2
        if potions[mid] < target {
            left = mid + 1  // 目标在右侧
        } else {
            right = mid     // 目标在左侧或当前mid
        }
    }
    return left
}
```
