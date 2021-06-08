package leetcode
func subsetsWithDup(nums []int) [][]int {

	result := make([][]int, 0)
	//bt78(nums, 0, []int{}, &result)
	list:=make([]int,0)
	var dfs func(int)
	dfs =func(pos int){
		result=append(result,list)
		for i:=pos;i<len(nums);i++{
			list=append(list,nums[i])
			dfs(i+1)
			//TODO

		}
	}
	dfs(0)
	return result

}
func bt90(nums []int, pos int, list []int, result *[][]int) {
	*result = append(*result, append([]int(nil),list...))
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		bt78(nums, i+1, list, result)
		list = list[:len(list)-1]
	}

}
