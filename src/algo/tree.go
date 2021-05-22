// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package algo

import (
	"fmt"
	"math/rand"
)

// A TreeNode is a binary tree with integer values.
type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *TreeNode {
	var t *TreeNode
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return &TreeNode{v, nil, nil}
	}
	if v < t.Val {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *TreeNode) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Val)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
