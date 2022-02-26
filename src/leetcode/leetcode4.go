package leetcode

//直接合并数组
func FindMedianSortedArrays_0(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	mergedNums := make([]int, 0, length)
	mid0, mid1 := 0, 0
	if length%2 != 0 {
		mid0 = length / 2
		mid1 = mid0
	} else {
		mid0 = length / 2
		mid1 = mid0 - 1
	}

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			mergedNums = append(mergedNums, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			mergedNums = append(mergedNums, nums2[j])
			j++
		}
	}
	mergedNums = append(mergedNums, nums1[i:]...)
	mergedNums = append(mergedNums, nums2[j:]...)
	midNum := float64(mergedNums[mid0]+mergedNums[mid1]) / 2
	return midNum
}

//移动两个指针
func FindMedianSortedArrays_1(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	midLeft, midRight := -1, -1
	p1, p2 := 0, 0
	for i := 0; i <= length/2; i++ {
		midLeft = midRight
		if p1 < len(nums1) && (p2 >= len(nums2) || nums1[p1] < nums2[p2]) {
			midRight = nums1[p1]
			p1++
		} else {
			midRight = nums2[p2]
			p2++
		}
	}
	if length%2 == 0 {
		return float64(midLeft+midRight) / 2
	} else {
		return float64(midRight)
	}
}
