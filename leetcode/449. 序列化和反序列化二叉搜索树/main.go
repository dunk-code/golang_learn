package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var ans string
	var prefixOrder func(node *TreeNode)
	prefixOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans += strconv.Itoa(node.Val) + " "
		if node.Left != nil {
			prefixOrder(node.Left)
		}
		if node.Right != nil {
			prefixOrder(node.Right)
		}
	}
	prefixOrder(root)
	return ans[:len(ans)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	strs := strings.Split(data, " ")
	nums := make([]int, len(strs))
	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}
	preOrder := make([]int, len(strs))
	for i, v := range nums {
		preOrder[i] = v
	}
	sort.Ints(nums)
	infixOrder := nums
	mp := map[int]int{}
	for i, v := range infixOrder {
		mp[v] = i
	}
	// fmt.Println("mp:", mp, "preOrder:", preOrder, "infixOrder:", infixOrder)
	var dfs func(preLeft, preRight, inLeft, inRight int) *TreeNode
	dfs = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		if preLeft > preRight {
			return nil
		}
		preRoot := preLeft
		inRoot := mp[preOrder[preRoot]]
		root := &TreeNode{Val: preOrder[preRoot]}
		sizeLeft := inRoot - inLeft
		root.Left = dfs(preLeft+1, preLeft+sizeLeft, inLeft, inRoot-1)
		root.Right = dfs(preLeft+sizeLeft+1, preRight, inRoot+1, inRight)
		// fmt.Println(root.Val, root.Left, root.Right)
		return root
	}
	return dfs(0, len(preOrder)-1, 0, len(infixOrder)-1)
}

func main() {
	root := &TreeNode{}
	codec := Constructor()
	data := codec.serialize(root)
	root = codec.deserialize(data)
	fmt.Println(root.Val)
	// newRoot := codec.deserialize(data)
	// fmt.Println(newRoot.Val, newRoot.Left.Val, newRoot.Right)
}
