package leetcode

func sum5863(path []int) int {
	sum := 0
	for _, v := range path {
		sum += v
	}
	return sum
}

func bt5863(nums, path []int, pos int, count *int) {
	if len(path) > 0 {
		if path[0] < sum5863(path[1:]) {
			return
		}
	}
	if len(path) == 4 {
		if path[0] == path[1]+path[2]+path[3] {
			*count++
		} else {
			return
		}
	}
	for i := pos; i < len(nums); i++ {
		path = append(path, nums[i])
		bt5863(nums, path, i+1, count)
		path = path[:len(path)-1]
	}
}

func CountQuadruplets(nums []int) int {
	count := 0
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	bt5863(nums, []int{}, 0, &count)
	return count
}
