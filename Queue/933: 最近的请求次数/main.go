package main

func main() {

}

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	if len(this.queue) == 0 {
		this.queue = append(this.queue, t)
		return len(this.queue)
	}

	left := t - 3000

	// 剔除不符合规则的请求
	for len(this.queue) > 0 && this.queue[0] < left {
		this.queue = this.queue[1:]
	}

	// 加入新请求
	this.queue = append(this.queue, t)

	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
