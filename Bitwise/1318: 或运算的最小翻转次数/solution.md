## 1318

### 题
给你三个正整数 a、b 和 c。

你可以对 a 和 b 的二进制表示进行位翻转操作，返回能够使按位或运算   a OR b == c  成立的最小翻转次数。

「位翻转操作」是指将一个数的二进制表示任何单个位上的 1 变成 0 或者 0 变成 1 。

### 题解
```go
func minFlips(a int, b int, c int) int {
	count := 0
	for a|b != c {
		// 提取出最低位
		aLow, bLow, cLow := a&1, b&1, c&1
		// 整体向右移动一位，并丢弃最低位，相当于检查下一位
		a, b, c = a>>1, b>>1, c>>1
		// 当前位已经相等，无需翻转
		if aLow|bLow == cLow {
			continue
		}
		// 如果 c 的当前位是 1，但 aLow|bLow 是 0，说明 a 或 b 至少有一个需要是 1。此时只需翻转其中一个
		if cLow == 1 {
			count++
		} else {
			// 如果 c 的当前位是 0，但 aLow|bLow 是 1，说明 a 和 b 的当前位必须都是 0
			if aLow == 1 {
				count++
			}
			if bLow == 1 {
				count++
			}
		}
	}
	return count
}
```