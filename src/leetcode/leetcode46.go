package leetcode

func bt46(path, nums []int, length int, result *[][]int) {
	if len(path) >= length {
		*result = append(*result, append([]int(nil), path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		newNums := append([]int(nil), nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		bt46(path, newNums, length, result)
		path = path[:len(path)-1]
	}
}

func in46(list []int, elem int) bool {
	for _, v := range list {
		if v == elem {
			return true
		}
	}
	return false
}

func bt46_2(nums []int, path []int, result *[][]int) {
	if len(path) >= len(nums) {
		*result = append(*result, append([]int(nil), path...))
	}
	for i := 0; i < len(nums); i++ {
		if in46(path, nums[i]) {
			continue
		}
		path = append(path, nums[i])
		bt46_2(nums, path, result)
		path = path[:len(path)-1]
	}

}

func bt46_3(nums,path []int,visited []bool,result *[][]int){
	if len(path)==len(nums){
		*result = append(*result, append([]int(nil), path...))
	}
	for i:=0;i<len(nums);i++{
		if visited[i]{
			continue
		}
		visited[i]=true
		path=append(path,nums[i])
		bt46_3(nums,path,visited,result)
		path=path[:len(path)-1]
		visited[i]=false
	}

}

// 传递去除元素的切片
func Permute0(nums []int) [][]int {
	result := make([][]int, 0)
	bt46([]int{}, nums, len(nums), &result)
	//bt46_2(nums, []int{}, &result)
	return result

}
// 检测元素是否存在，传入数组不能有重复元素
func Permute1(nums []int) [][]int {
	result := make([][]int, 0)
	//bt46([]int{}, nums, len(nums), &result)
	bt46_2(nums, []int{}, &result)
	return result

}
// visited 数组
func Permute2(nums []int) [][]int {
	result := make([][]int, 0)
	visited:=make([]bool,len(nums))
	bt46_3(nums,[]int{},visited,&result)
	return result
	
}
