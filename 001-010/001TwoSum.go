// 1. Two Sum - https://leetcode.com/problems/two-sum/
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

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

	for i, v := range nums {
		complement := target - v
		if j, ok := m[compliment]; ok {
			return []int{j, i} // found answer
		}
		m[v] = i
	}
	return nil // no answer
}