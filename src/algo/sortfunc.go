package algo

//var a =[]int{9,3,41,7,98,621,1,4}


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
func MergeSort(arr []int)[]int{
	if len(arr)<=1{
		return arr
	}
	mid:=len(arr)/2
	left:=MergeSort(arr[:mid])
	right:=MergeSort(arr[mid:])
	result:=merge(left,right)
	return result
}

func QuickSort0(arr []int,leftEnd,rightEnd int){
	if leftEnd>rightEnd{
		return 
	}
	left:=leftEnd
	right:=rightEnd
	pivot:=arr[left]
	for left<right{
		for left<right && arr[right]>=pivot{
			right-=1
		}
		if left<right{
			arr[left]=arr[right]
		}
		for left<right && arr[left]<=pivot{
			left+=1
		}
		if left<right{
			arr[right]=arr[left]
		}
		if left>=right{
			arr[left]=pivot
		}
	}
	QuickSort0(arr,leftEnd,right-1)
	QuickSort0(arr,left+1,rightEnd)
}

