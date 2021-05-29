package leetcode

import "fmt"
var maxDiameter int=0

func dfs543(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDiameter:=dfs543(root.Left)
	rDiameter:=dfs543(root.Right)
	if lDiameter+rDiameter>maxDiameter{
		maxDiameter=lDiameter+rDiameter
		fmt.Println("md修改了",maxDiameter)
	}
	if lDiameter>=rDiameter{
		return lDiameter+1
	}else{
		return rDiameter+1
	}
}
func DiameterOfBinaryTree(root *TreeNode) int {
	dfs543(root)
	return maxDiameter


}
