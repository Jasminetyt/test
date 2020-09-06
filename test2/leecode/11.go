package main

import "fmt"

func main() {
	var height=[]int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}
func maxArea(height []int) int {
	var front,back=0,len(height)-1
	var max,area int
	for front<back{
		temp:=back-front
		if height[front]>height[back] {
			area=temp*height[back]
			back--
		}else {
				area=temp*height[front]
				front++
		}
		if area>max {
			max=area
		}
	}
	return max
}