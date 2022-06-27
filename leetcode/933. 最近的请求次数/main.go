package main

import "fmt"

type RecentCounter struct {
	requestNums []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requestNums: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.requestNums = append(this.requestNums, t)
	leftLmt := t - 3000
	idx := binarySearch(this.requestNums, leftLmt)
	this.requestNums = this.requestNums[idx:]
	return len(this.requestNums)
}

func binarySearch(nums []int, tar int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= tar {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	arr := []int{1, 5, 7, 9, 19, 30}
	search := binarySearch(arr, 5)
	arr = arr[search:]
	fmt.Println(arr)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
