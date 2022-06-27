package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) (ans *TreeNode) {
	if p.Right != nil {
		ans = p.Right
		for ans.Left != nil {
			ans = ans.Left
		}
		return ans
	}
	tmp := root
	for tmp != nil {
		if tmp.Val > p.Val {
			ans = tmp
			tmp = tmp.Left
		} else {
			tmp = tmp.Right
		}
	}
	return
}

func main() {

}
