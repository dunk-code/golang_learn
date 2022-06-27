package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for i := 0; queue != nil; i++ {
		tmp := queue
		queue = nil
		row := make([]int, 0, len(tmp))
		for _, cur := range tmp {
			row = append(row, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		if i%2 == 1 {
			l, r := 0, len(row)-1
			for l <= r {
				row[r], row[l] = row[l], row[r]
				l++
				r--
			}
		}
		ans = append(ans, row)
	}
	return ans
}

func main() {

}
