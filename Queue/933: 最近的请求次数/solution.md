## 933

### 题
写一个 RecentCounter 类来计算特定时间范围内最近的请求。

请你实现 RecentCounter 类：

RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
保证 每次对 ping 的调用都使用比之前更大的 t 值。

### 题解
```go
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
```
维护一个队列，本题想要的结果就是最终的队列长度。超出时间范围的请求出队、新请求进队。
