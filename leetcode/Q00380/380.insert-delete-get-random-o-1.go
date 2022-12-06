/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */
package q00380

import "math/rand"

// @lc code=start
type RandomizedSet struct {
    Item map[int]int
	CurrArr []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
		Item: map[int]int{},
		CurrArr: make([]int, 0),
	}
}


func (this *RandomizedSet) Insert(val int) bool {
    _, ok := this.Item[val]; 
	if !ok {
		this.Item[val] = len(this.CurrArr)
		this.CurrArr = append(this.CurrArr, val)
	} 
	return !ok
}


func (this *RandomizedSet) Remove(val int) bool {
    index, exist := this.Item[val]
	if exist {
		delete(this.Item, val)
		lastIdx := len(this.CurrArr) - 1
		if index != lastIdx {
			tmp := this.CurrArr[lastIdx]
			this.Item[tmp] = index
			this.CurrArr[index] = tmp
		}
		this.CurrArr = this.CurrArr[:lastIdx]
	}
	return exist
}


func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.CurrArr))
	return this.CurrArr[i]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

