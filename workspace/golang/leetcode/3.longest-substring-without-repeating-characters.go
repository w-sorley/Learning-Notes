/*
 * @Author: wang shouli
 * @LastEditors: wang shouli
 */
/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (29.16%)
 * Likes:    7840
 * Dislikes: 467
 * Total Accepted:    1.3M
 * Total Submissions: 4.5M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 *
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var max = 0
	var charMap = make(map[rune]int)
	var startIndex = 0
	for index, ch := range s {
		if _, ok := charMap[ch]; ok {
			startIndex = Max(startIndex, charMap[ch])
		}
		max = Max(max, index+1-startIndex)
		charMap[ch] = index + 1
	}
	return Max(max, len(s)-startIndex)
}

// Max  max
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}