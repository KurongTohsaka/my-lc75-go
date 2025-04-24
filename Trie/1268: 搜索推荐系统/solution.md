## 1268

### 题
给你一个产品数组 products 和一个字符串 searchWord ，products  数组中每个产品都是一个字符串。

请你设计一个推荐系统，在依次输入单词 searchWord 的每一个字母后，推荐 products 数组中前缀与 searchWord 相同的最多三个产品。如果前缀相同的可推荐产品超过三个，请按字典序返回最小的三个。

请你以二维列表的形式，返回在输入 searchWord 每个字母后相应的推荐产品的列表。

### 题解
```go
func suggestedProducts(products []string, searchWord string) [][]string {
	tree := newTrieNode()

	for _, word := range products {
		tree.insert(word)
	}

	res := tree.search(searchWord)

	return res
}

type TrieNode struct {
	words    []string
	children [26]*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{}
}

func (t *TrieNode) insert(word string) {
	root := t

	for _, c := range word {
		idx := int(c - 'a')
		if root.children[idx] == nil {
			root.children[idx] = newTrieNode()
		}
		root = root.children[idx]
		root.words = append(root.words, word)
		sort.Strings(root.words)
		if len(root.words) > 3 {
			root.words = root.words[:3]
		}
	}
}

func (t *TrieNode) search(word string) [][]string {
	var result [][]string
	root := t
	flag := false // 标记是否已经遇到无法匹配的前缀

	// 最终返回的 result 是一个二维数组，其中每个元素对应搜索词的一个前缀的推荐结果
	// 例如：搜索词 "mouse" 会返回 5 个数组，分别对应 "m", "mo", "mou", "mous", "mouse" 的前缀推荐
	for _, c := range word {
		if flag {
			result = append(result, make([]string, 0))
		} else {
			idx := int(c - 'a')
			if root.children[idx] == nil {
				result = append(result, make([]string, 0))
				flag = true
			} else {
				root = root.children[idx]
				result = append(result, root.words)
			}
		}
	}
	return result
}
```
手搓前缀树解法。