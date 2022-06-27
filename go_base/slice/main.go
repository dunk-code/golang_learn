package main

import "fmt"

func printArray(arr []int) {
	fmt.Printf("printArray >>> enter arr addr : %p\n", &arr)
	fmt.Println(arr)
}
func changeArr(arr []int) {
	arr[0] = -1
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Printf("nums addr:%p\n", &nums)
	fmt.Printf("nums type %T\n", nums)
	printArray(nums)
	changeArr(nums)
	printArray(nums)
}
