## 735
### 题
给定一个整数数组 asteroids，表示在同一行的小行星。数组中小行星的索引表示它们在空间中的相对位置。

对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。

找出碰撞后剩下的所有小行星。碰撞规则：两个小行星相互碰撞，较小的小行星会爆炸。如果两颗小行星大小相同，则两颗小行星都会爆炸。两颗移动方向相同的小行星，永远不会发生碰撞。

### 题解
```go
func asteroidCollision(asteroids []int) []int {
	var stack []int

	for _, val := range asteroids {
		// 只有栈顶元素向右（>0）、val 向左（<0）才有可能碰撞
		// 需要循环进行判断，直到栈内元素可碰撞的全部毁掉
		for len(stack) > 0 && val < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top > -val {
				val = 0
				break
			} else if top == -val {
				stack = stack[:len(stack)-1]
				val = 0
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}

		if val != 0 {
			stack = append(stack, val)
		}
	}

	return stack
}
```
只有当栈顶向右（正）且当前行星向左（负）时才处理碰撞，使用循环处理可能的多次碰撞，直到当前行星被摧毁或无法继续碰撞。
