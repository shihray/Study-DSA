/*
 * @lc app=leetcode id=705 lang=golang
 *
 * [705] Design HashSet
 */

// @lc code=start
type MyHashSet struct {
    List [10000][]int
	Size int
}


func Constructor() MyHashSet {
	return MyHashSet{Size: 10000, List: [10000][]int{}}
}


func (this *MyHashSet) Add(key int)  {
    hash := key % 10000

	this.List[hash] = append(this.List[hash], key)
}


func (this *MyHashSet) Remove(key int)  {
    hash := key % 10000

	for idx, val := range this.List[hash] {
		if val == key {
			this.List[hash][idx] = -1
		}
	}
}


func (this *MyHashSet) Contains(key int) bool {
    hash := key % 10000

	for _, v := range this.List[hash] {
		if v == key {
			return true
		}
	}
	return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

