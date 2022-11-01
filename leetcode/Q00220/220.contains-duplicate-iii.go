/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 */
package q00220

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {

	bucket := make(map[int]int)
	width := valueDiff + 1

	for idx, num := range nums {
		var bucketIdx int

		if num >= 0 {
			bucketIdx = num / width
		} else {
			bucketIdx = (num / width) - 1
		}

		if _, ok := bucket[bucketIdx]; ok {
			return true
		} else if _, ok := bucket[bucketIdx+1]; ok && (Abs(num-bucket[bucketIdx+1]) < width) {
			return true
		} else if _, ok := bucket[bucketIdx-1]; ok && (Abs(num-bucket[bucketIdx-1]) < width) {
			return true
		}

		bucket[bucketIdx] = num
		if idx >= indexDiff {
			delete(bucket, (nums[idx-indexDiff] / width))
		}
	}
	return false
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

// @lc code=end
