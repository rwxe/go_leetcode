package leetcode

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
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
