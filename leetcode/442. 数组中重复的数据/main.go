package main

/*func findDuplicates(nums []int) (ans []int) {
	mp := make([]int, len(nums)+1)
	for _, num := range nums {
		if mp[num] == 0 {
			mp[num] = 1
		} else {
			ans = append(ans, num)
		}
	}
	return ans
}*/

func findDuplicates(nums []int) (ans []int) {
	for i := range nums {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num != i+1 {
			ans = append(ans, num)
		}
	}
	return
}
func main() {

}
