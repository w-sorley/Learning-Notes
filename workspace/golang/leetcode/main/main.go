/*
 * @Author: wang shouli
 * @LastEditors: wang shouli
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello golang !!")
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var max = 0
	var charMap = make(map[rune]int)
	var startIndex = 0
	for index, ch := range s {
		if _, ok := charMap[ch]; ok {
			startIndex = charMap[ch]
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
