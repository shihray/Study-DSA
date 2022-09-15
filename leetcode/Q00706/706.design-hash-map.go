/*
 * @lc app=leetcode id=706 lang=golang
 *
 * [706] Design HashMap
 */
package q00706

// @lc code=start
type MyHashMap struct {
    Set []int
	Size int
}


func Constructor() MyHashMap {
	size := 1000000
	resp := MyHashMap{
		Set: make([]int, size),
		Size: size,
	}

	for i := 0; i < size; i++ {
		resp.Set[i] = -1
	}

    return resp
}


func (this *MyHashMap) Put(key int, value int)  {
    hash := key % this.Size

	this.Set[hash] = value
}


func (this *MyHashMap) Get(key int) int {
    hash := key % this.Size

	return this.Set[hash]
}


func (this *MyHashMap) Remove(key int)  {
    hash := key % this.Size
	this.Set[hash] = -1
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

