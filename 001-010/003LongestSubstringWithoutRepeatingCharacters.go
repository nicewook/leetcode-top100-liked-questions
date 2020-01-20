// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
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

/*
고려사항
1. 빈 string 인 경우는?
2. nil string 인 경우는?
3. 본적이 있는 녀석이면 이전에 본거 다음 녀석이 start index 가 된다.
	- 그 전까지의 길이를 저장해둔다.
*/

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
		lastSeen[c] = i // 무조건 최신으로, 마지막에 본 위치로 업데이트
	}
	lastLen := len(s) - start
	if max < lastLen {
		max = lastLen
	}
	return max
}