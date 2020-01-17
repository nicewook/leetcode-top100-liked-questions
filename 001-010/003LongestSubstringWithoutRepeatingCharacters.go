// 3. Longest Substring Without Repeating Characters
// Given a string, find the length of the longest substring without repeating characters.

// Example 1:

// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	var max, start int
	lastSeen := make(map[rune]int) // key: charactor, value: index
	for i, c := range s {
		if index, ok := lastSeen[c]; ok && index >= start {
			curLen := i - start
			if max < curLen {
				max = curLen
			}
			start = index + 1
		}
		lastSeen[c] = i
	}
	lastLen := len(s) - start
	if max < lastLen {
		max = lastLen
	}
	return max
}