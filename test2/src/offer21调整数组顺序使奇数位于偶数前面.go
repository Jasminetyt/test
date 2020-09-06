package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6}
	reorder(nums)
	fmt.Println(nums)
}
func reorder(nums []int)  {
	if nums==nil {
		return
	}
	var front ,back=0,len(nums)-1
	for front<back {
		for isEvent(nums[front]) { //若为奇数则后移
			front++
		}
		for ! isEvent(nums[back])  { //若为偶数则前移
			back--
		}
		if front < back {
			nums[front],nums[back]=nums[back],nums[front]
		}
	}
}

func isEvent(num int) bool {
	return num & 0x1 == 1 //num为奇数返回true
}