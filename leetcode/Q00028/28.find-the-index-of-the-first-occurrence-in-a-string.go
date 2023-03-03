/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

package q00028

func strStr(haystack string, needle string) int {

	nlen := len(needle)
	hlen := len(haystack)

	for i := 0; i < hlen; i++ {

		nI := 0

		if haystack[i] == needle[nI] {
			if i+nlen > hlen {
				return -1
			}
			if haystack[i:i+nlen] == needle {
				return i
			}
		}
	}
	return -1
}
