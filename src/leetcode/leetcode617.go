package leetcode

func MergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1==nil{return root2}
	if root2==nil{return root1}
	root1.Val+=root2.Val
	root1.Left=MergeTrees1(root1.Left,root2.Left)
	root1.Right=MergeTrees1(root1.Right,root2.Right)
	return root1
}

func MergeTrees0(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil || root2 == nil {
		if root1 == nil {
			return root2
		} else if root2 == nil {
			return root1
		}
	}
	rootMerged:=&TreeNode{Val:root1.Val+root2.Val}
	queue1:=make([]*TreeNode,0)
	queue2:=make([]*TreeNode,0)
	queueMerged:=make([]*TreeNode,0)
	queue1=append(queue1,root1)
	queue2=append(queue2,root2)
	queueMerged=append(queueMerged,rootMerged)
	for len(queue1)>0 && len(queue2)>0{
		nodeMerged:=queueMerged[0]
		node1:=queue1[0]
		node2:=queue2[0]
		queueMerged=queueMerged[1:]
		queue1=queue1[1:]
		queue2=queue2[1:]
		left1,right1:=node1.Left,node1.Right
		left2,right2:=node2.Left,node2.Right

		if left1!=nil && left2!=nil{
			nodeMerged.Left=&TreeNode{Val:left1.Val+left2.Val}
			queueMerged=append(queueMerged,nodeMerged.Left)
			queue1=append(queue1,left1)
			queue2=append(queue2,left2)
		}else if left1!=nil && left2==nil{
			nodeMerged.Left=left1
		} else if left1==nil && left2!=nil{
			nodeMerged.Left=left2
		}
		if right1!=nil && right2!=nil{
			nodeMerged.Right=&TreeNode{Val:right1.Val+right2.Val}
			queueMerged=append(queueMerged,nodeMerged.Right)
			queue1=append(queue1,right1)
			queue2=append(queue2,right2)
		}else if right1!=nil && right2==nil{
			nodeMerged.Right=right1
		} else if right1==nil && right2!=nil{
			nodeMerged.Right=right2
		}
	}
	return rootMerged
}
