package main

func maxRotateFunction(nums []int) int {
	numSum := 0
	for _, v := range nums {
		numSum += v
	}
	f := 0
	for i, v := range nums {
		f += i * v
	}
	ans := f
	for i := len(nums) - 1; i > 0; i-- {
		f += numSum - len(nums)*nums[i]
		ans = max(ans, f)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	print(maxRotateFunction([]int{4, 3, 2, 6}))
}
