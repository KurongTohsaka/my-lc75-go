package main

func main() {

}

func rob(nums []int) int {
	n := len(nums)
	cache := make([]int, n+1)
	cache[0], cache[1] = 0, nums[0]
	for i := 2; i <= n; i++ {
		// cache[i] 表示到第i个房屋为止所能偷取的最大金额
		// cache[i-1] 表示相邻房屋的最大值，不加 nums[i-1] 因为相邻房屋不能连续偷取
		// cache[i-2]+nums[i-1] 相隔的房屋的最大值加上当前房屋的值
		cache[i] = max(cache[i-1], cache[i-2]+nums[i-1])
	}

	return cache[n]
}
