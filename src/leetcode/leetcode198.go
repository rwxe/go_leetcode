package leetcode
func rob(nums []int) int {
	max:=func(a,b int)int{if a>b{return a}else{return b}}
	w:=make([]int,len(nums))
	for i:=range w{
		if i==0{
			w[i]=nums[0]
		}else if i==1{
			w[i]=max(nums[1],nums[0])
		}else if i>1{
			w[i]=max(w[i-2]+nums[i],w[i-1])
		}
	}
	return w[len(w)-1]

}
