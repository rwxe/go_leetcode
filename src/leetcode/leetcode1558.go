package leetcode

func minOperations(nums []int) int {
	srcNums := make([]int, len(nums))
	result := 0
	for i := range nums {
		queue := make([]int, 0)
		queue = append(queue, srcNums[i])
		count := 0
		for len(queue) != 0 {
			flag := false
			width := len(queue)
			for j := 0; j < width; j++ {
				node := queue[0]
				queue = queue[1:]
				if node == nums[i] {
					flag = true
					break
				}
				if node < nums[i] {
					queue = append(queue, node*2)
					queue = append(queue, node+1)
				}
			}
			if flag {
				break
			}
			count++
		}
		result += count
	}
	return result
}
