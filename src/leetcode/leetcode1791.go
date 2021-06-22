package leetcode


func allTrue(elem ...bool)bool{
	for _,v:=range elem{
		if !v{
			return false
		}
	}
	return true
}

// 会爆内存
func FindCenter(edges [][]int) int {
	adjacentMatrix:=make([][]bool,len(edges)+1)
	for i:=range adjacentMatrix{
		adjacentMatrix[i]=make([]bool,len(edges)+1)
		adjacentMatrix[i][i]=true
	}
	for _,edge:=range edges{
		adjacentMatrix[edge[0]-1][edge[1]-1]=true
		adjacentMatrix[edge[1]-1][edge[0]-1]=true
	}
	for i,v:=range adjacentMatrix{
		if allTrue(v...){
			return i+1
		}


	}
	return 0

}

func FindCenter2(edges [][]int) int {
	if edges[0][0]==edges[1][0]{
		return edges[0][0]
	}else if edges[0][0]==edges[1][1]{
		return edges[0][0]
	}else if edges[0][1]==edges[1][0]{
		return edges[0][1]
	}else  if edges[0][1]==edges[1][1]{
		return edges[0][1]
	}else{
		return edges[0][1]
	}
	
}
