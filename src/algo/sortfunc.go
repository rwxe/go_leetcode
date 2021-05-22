package algo

//var a =[]int{9,3,41,7,98,621,1,4}

func MergeSort(nums []int)[]int{
	if len(nums)<=1{
		return nums
	}
	mid:=len(nums)/2
	left:=MergeSort(nums[:mid])
	right:=MergeSort(nums[mid:])
	result:=merge(left,right)
	return result
}

func merge(left,right []int)[]int{
	l:=0
	r:=0
	result:=make([]int,0)
	for l<len(left) && r<len(right){
		if left[l]<right[r]{
			result=append(result,left[l])
			l++
		}else{
			result=append(result,right[r])
			r++
		}
	}
	result=append(result,left[l:]...)
	result=append(result,right[r:]...)
	return result
}
