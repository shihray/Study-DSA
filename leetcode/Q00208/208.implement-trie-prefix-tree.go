/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */
package q00208

// @lc code=start
type Trie struct {
	child [26]*Trie
	end bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, c := range word {
		idx := c - 'a'
		if curr.child[idx] == nil {
			curr.child[idx] = &Trie{}
		}
		curr = curr.child[idx]
	}
	curr.end = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, c := range word {
		idx := c - 'a'
		if curr.child[idx] == nil {
			return false
		}
		curr = curr.child[idx]
	}
	return curr.end
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
    for _, c := range prefix {
        idx := c - 'a'
        if curr.child[idx] == nil {
            return false
        }
        curr = curr.child[idx]
    }
    return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
