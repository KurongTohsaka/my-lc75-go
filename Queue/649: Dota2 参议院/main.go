package main

func main() {

}

func predictPartyVictory(senate string) string {
	var rQueue, dQueue []int

	for i, ch := range senate {
		if ch == 'R' {
			rQueue = append(rQueue, i)
		} else {
			dQueue = append(dQueue, i)
		}
	}

	for len(rQueue) > 0 && len(dQueue) > 0 {
		// 每个议员既可以淘汰别人，也可以被任何人淘汰。所以这里加上len(senate)是在降低已经淘汰过别人的议员权重
		if rQueue[0] < dQueue[0] {
			rQueue = append(rQueue, rQueue[0]+len(senate))
		} else {
			dQueue = append(dQueue, dQueue[0]+len(senate))
		}
		// 这里淘汰的两个人，一个是降低权重到队列后方的、一个是被淘汰的
		rQueue = rQueue[1:]
		dQueue = dQueue[1:]
	}

	if len(dQueue) > 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}
