package leetcode

func Merge88_2(nums1 []int, m int, nums2 []int, n int) {
	for i, j, k := m-1, n-1, m+n-1; i >= 0 || j >= 0;k-- {
		if i<0{
			nums1[k]=nums2[j]
			j--
		}else if j<0{
			nums1[k]=nums1[i]
			i--
		}else if nums1[i]>nums2[j]{
			nums1[k]=nums1[i]
			i--
		}else{
			nums1[k]=nums2[j]
			j--
		}
	}

}
func Merge88(nums1 []int, m int, nums2 []int, n int) {
	//nums1 [0:m]
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]

	}

	for i, j, k := n, 0, 0; i < m+n || j < n; k++ {
		if i >= m+n {
			nums1[k] = nums2[j]
			j++
		} else if j >= n {
			nums1[k] = nums1[i]
			i++
		} else if nums1[i] <= nums2[j] {
			nums1[k] = nums1[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}
}
