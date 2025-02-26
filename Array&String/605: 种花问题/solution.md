## 605

### 题
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。

```go
package main

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n1 := 1
	n2 := 2
	println(canPlaceFlowers(flowerbed, n1))
	println(canPlaceFlowers(flowerbed, n2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)

	// 边界
	if length == 1 {
		if flowerbed[0] == 0 {
			n -= 1
		}
	} else {
		for i := 0; i < length; i++ {
			if flowerbed[i] == 0 && ((i == 0 && flowerbed[i+1] == 0) || (i == length-1 && flowerbed[i-1] == 0) || (i > 0 && i < length-1 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0)) {
				flowerbed[i] = 1
				n -= 1
			}
		}
	}
	return n <= 0
}

```
当花坛只有一个位置时，如果该位置为空则直接种一朵花。
当花坛长度大于1时，遍历数组：
对于第一个位置（i==0），只需检查右边是否为空；
对于最后一个位置（i==length-1），只需检查左边是否为空；
对于中间位置，则要求左右均为空。
每种一朵花就将 n 减1，若 n 小于等于0则提前返回 true。