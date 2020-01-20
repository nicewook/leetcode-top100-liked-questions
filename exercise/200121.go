// 1. Two Sum
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1]

/*
고려사항
1. 해시테이블을 쓰면 속도가 빨라진다.
	- 해시의 키는 찾는 보수(complement), 값은 인덱스이다.
2. 바로 모든 값들이 해시되어 있을 필요는 없다.
	- 현재까지 해시되어 있는 값에서 없고, 뒤에 있다면 뒤에서 앞의 녀석을 찾을 것이다.package exercise
3. 이왕이면, j, i 로 리턴해준 게 포인트
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

// 2. Add Two Numbers
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain
// a single digit. Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example:

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

/*
고려사항
1. 두 linked list 가 길이가 다른 경우를 고려하자
2. 저장해야 할 값이 있을때만 linked list를 생성해야 한다. 
3. 처음 시작하는 녀석을 떼어 내고 .Next 를 리턴하는게 포인트
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(a *ListNode, b *ListNode) *ListNode {
	head := new(ListNode)
	tail : head
	var num int

	for a != nil || b != nil {
		num = num / 10 // carry
		if a != nil {
			num += a.Val
			a = a.Next
		}
		if b != nil {
			num += a.Val
			a = a.Next
		}
		tail.Next = &ListNode{num%10, nil}
		tail = tail.Next
	}
	if num/10 == 1 {
		tail.Next = &ListNode{1, nil}
	}
	return head.Next
}

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

/*
고려사항
1. 빈 string 인 경우는?
2. nil string 인 경우는?
3. 본적이 있는 녀석이면 이전에 본거 다음 녀석이 start index 가 된다.
	- 그 전까지의 길이를 저장해둔다.
*/

func lengthOfLongestSubstring(s string) int {
	var max, start int
	lastSeen := make(map(rune)int)
	for i, c := range s {
		if idx, ok := lastSeen[r]; ok && ldx >= start {  
			curLen := i - start
			if max < curLen {
				max = curLen
			}
			start = idx + 1
		}
		lstSeen[c] = i
	}
	lastLen := len(s) - start
	if max < lastLen {
		max = lastLen
	}

	return max
}















