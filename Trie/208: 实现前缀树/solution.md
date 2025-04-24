## 208

### 题
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

### 题解
```go
type Node struct {
	son [26]*Node
	end bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{&Node{}}
}

func (t Trie) Insert(word string) {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil {
			cur.son[c] = &Node{}
		}
		cur = cur.son[c]
	}
	cur.end = true
}

func (t Trie) find(word string) int {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil {
			return 0
		}
		cur = cur.son[c]
	}
	if cur.end {
		return 2
	}
	return 1
}

func (t Trie) Search(word string) bool {
	return t.find(word) == 2
}

func (t Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != 0
}
```
手搓前缀树，面试很少考，原理可以看看。
