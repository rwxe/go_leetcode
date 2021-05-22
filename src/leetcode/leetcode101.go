package leetcode


func Walk(t *TreeNode, turnLeft bool, ch chan *TreeNode) {
	queue:=make([]*TreeNode,0)
	queue=append(queue,t)
	for len(queue)>0{
		levelLen:=len(queue)
		for i:=0;i<levelLen;i++{
			levelNode:=queue[0]
			queue=queue[1:]
			ch<-levelNode
			if levelNode!=nil{
				if turnLeft{
					queue=append(queue,levelNode.Left)
					queue=append(queue,levelNode.Right)
				}else{
					queue=append(queue,levelNode.Right)
					queue=append(queue,levelNode.Left)

				}
			}


		}

	}
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *TreeNode) bool {
	ch1 := make(chan *TreeNode)
	ch2 := make(chan *TreeNode)
	go Walk(t1, true, ch1)
	go Walk(t2, false, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 && ok2 {
			//两个信道都在开启状态
			if v1 ==nil || v2==nil{
				if v1!=v2{
					return false
				}
			}else if v1.Val != v2.Val {
				return false
			}
		} else if ok1 || ok2 {
			//一个信道已关闭，说明两棵树节点数不同
			return false
		} else {
			//两个信道同时关闭，说明两棵树节点数相同
			break
		}
	}
	return true
}
func check(left,right *TreeNode) bool{
    if left==nil && right==nil{
        return true
    }else if left==nil || right==nil{
        return false
    }
    return left.Val==right.Val && check(left.Left,right.Right) && check(left.Right,right.Left)
}
func IsSymmetricRef(root *TreeNode) bool {
	return check(root.Left,root.Right)
}
func IsSymmetricGR(root *TreeNode) bool {
	if root==nil{
		return true
	}else if root.Left.Val != root.Right.Val{
		return false
	}
	return Same(root.Left,root.Right)

}
func isSymmetric(root *TreeNode) bool {
    queueLeft:=make([]*TreeNode,0)
    queueRight:=make([]*TreeNode,0)
            queueLeft=append(queueLeft,root)
        queueRight=append(queueRight,root)
    for len(queueLeft)>0 && len(queueRight)>0{
        left:=queueLeft[0]
        right:=queueRight[0]
        queueLeft=queueLeft[1:]
        queueRight=queueRight[1:]
        if left==nil && right==nil{
            continue
        }
        if left==nil || right==nil{
            return false
        }
        if left.Val!=right.Val{
            return false
        }
        queueLeft=append(queueLeft,left.Left)
        queueRight=append(queueRight,right.Right)
        queueLeft=append(queueLeft,left.Right)
        queueRight=append(queueRight,right.Left)

    }
    return true


}
