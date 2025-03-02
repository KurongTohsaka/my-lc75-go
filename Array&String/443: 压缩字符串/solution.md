## 443

### 题
给你一个字符数组 chars ，请使用下述算法压缩：

从一个空字符串 s 开始。对于 chars 中的每组 连续重复字符 ：

如果这一组长度为 1 ，则将字符追加到 s 中。
否则，需要向 s 追加字符，后跟这一组的长度。
压缩后得到的字符串 s 不应该直接返回 ，需要转储到字符数组 chars 中。需要注意的是，如果组长度为 10 或 10 以上，则在 chars 数组中会被拆分为多个字符。

请在 修改完输入数组后 ，返回该数组的新长度。

你必须设计并实现一个只使用常量额外空间的算法来解决此问题。

```go
func compress(chars []byte) int {
	n := len(chars)
	write, conStart := 0, 0
	for read := 0; read < n; read++ {
		if read == n-1 || chars[read] != chars[read+1] {
			chars[write] = chars[read]
			write++
			count := read - conStart + 1
			if count > 1 {
				cntStr := strconv.Itoa(count)
				for _, c := range []byte(cntStr) {
					chars[write] = c
					write++
				}
			}
			conStart = read + 1
		}
	}
	return write
}
```

### 题解
该题使用双指针解题。
需要注意的是本题不能使用额外的空间，而且必须原地修改数组。
上面代码中的 `read` 、 `write` 、 `conStart` 分别是读指针、新的压缩字符串位置（初值为0是因为原地修改数组）、连续数组起始位置。 以下是执行流程，总体来说还是比较复杂的：
- `read == n-1 || chars[read] != chars[read+1]` 这一段条件表示连续序列已经结束，之后就是对这一段连续序列的处理。
- 先是将当前字符写入压缩字符串 `chars[write] = chars[read]`，然后是求这一段连续序列长度 `count = read - conStart + 1`，之后 `conStart` 更新新的序列起始位置。
- 对于 `count` ，先把 int 转换为字符串，然后把每位依次写入新的压缩字符串对应位置，同时更新 `write` 。
- 最后返回压缩串长度 `write` 。 

