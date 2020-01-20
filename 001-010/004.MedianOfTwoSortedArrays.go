// 4. Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/

// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.

// Example 1:
// nums1 = [1, 3]
// nums2 = [2]
// The median is 2.0

// Example 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// The median is (2 + 3)/2 = 2.5

/*
고민
1. 그냥 하나의 array 로 만든 다음에 가장 가운데 값, 혹은 가운데 두 값의 합을 2로 나눈값을 구하면 된다.
2. 하지만 O(log (m+n)) 이어야 한다.
*/

