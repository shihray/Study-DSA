/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */
package q00912

// @lc code=start
func sortArray(nums []int) []int {
    quicksort(nums, 0, len(nums)-1)
    return nums
}

func quicksort(arr []int, left, right int) {
    if left >= right { return }
    pivot := partition(arr, left, right)
    quicksort(arr, left, pivot-1)
    quicksort(arr, pivot+1, right)
}

func partition(arr []int, left, right int) int {
    pivot := right
    for left < right {
        for left < right && arr[left] <= arr[pivot] { left++ }
        for left < right && arr[right] >= arr[pivot] { right-- }
        if left < right { arr[left], arr[right] = arr[right], arr[left] }
    }
    arr[pivot], arr[right] = arr[right], arr[pivot] // set pivot at the correct position
    return right
}
// @lc code=end

