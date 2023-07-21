package leetcode

import "math"

//二分O(m,n)
//分隔线的性质
// nums1L<=nums2R && nums2L<=nums1R
// 对较短数组进行二分，保证较长数组不会越界
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var numsM, numsN []int
	if len(nums1) <= len(nums2) {
		numsM = nums1
		numsN = nums2
	} else {
		numsM = nums2
		numsN = nums1
	}
	m, n := len(numsM), len(numsN)
	totalLeft := (m + n + 1) / 2
	left, right := 0, m
	for left < right {
		i := (left + right + 1) / 2
		j := totalLeft - i
		//numsMLeftMax<=numsNRightMin
		//numsNLeftMax<=numsMRightMin
		if numsM[i-1] > numsN[j] {
			right = i - 1
		} else {
			left = i
		}
	}

	i := left
	j := totalLeft - i
	var numsMILeft, numsMIRight, numsNJLeft, numsNJRight int
	if i == 0 {
		numsMILeft = math.MinInt64
	} else {
		numsMILeft = numsM[i-1]
	}
	if i == m {
		numsMIRight = math.MaxInt64
	} else {
		numsMIRight = numsM[i]
	}
	if j == 0 {
		numsNJLeft = math.MinInt64
	} else {
		numsNJLeft = numsN[j-1]
	}
	if j == n {
		numsNJRight = math.MaxInt64
	} else {
		numsNJRight = numsN[j]
	}

	if (m+n)%2 == 0 {
		return (math.Max(float64(numsMILeft), float64(numsNJLeft)) +
			math.Min(float64(numsMIRight), float64(numsNJRight))) / 2
	} else {
		return math.Max(float64(numsMILeft), float64(numsNJLeft))
	}

}

// 双指针 O(m+n)
func findMedianSortedArrays_0(nums1 []int, nums2 []int) float64 {
	var mid0, mid1 int
	var medianSum int
	if sumLen := len(nums1) + len(nums2); sumLen == 0 {
		return float64(0)
	} else if sumLen%2 == 0 {
		mid1 = sumLen / 2
		mid0 = mid1 - 1
	} else {
		mid1 = sumLen / 2
		mid0 = mid1
	}

	for index, p1, p2 := 0, 0, 0; p1 < len(nums1) || p2 < len(nums2); index++ {
		var currentIsP1 bool

		if p1 < len(nums1) && p2 < len(nums2) && nums1[p1] < nums2[p2] {
			currentIsP1 = true
		} else if p1 >= len(nums1) {
			currentIsP1 = false
		} else if p2 >= len(nums2) {
			currentIsP1 = true
		} else { //p1 < len(nums1) && p2 < len(nums2) && nums1[p1] >= nums2[p2]
			currentIsP1 = false
		}

		if currentIsP1 {
			if index == mid0 {
				medianSum += nums1[p1]
			}
			if index == mid1 {
				medianSum += nums1[p1]
			}
			p1++
		} else {
			if index == mid0 {
				medianSum += nums2[p2]
			}
			if index == mid1 {
				medianSum += nums2[p2]
			}
			p2++
		}
	}
	return float64(medianSum) / 2
}
