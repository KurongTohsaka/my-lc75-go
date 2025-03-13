package main

func main() {

}

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
