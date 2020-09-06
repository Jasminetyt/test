package main

import "fmt"

func main()  {
	 arr := []int{0,0,1,2,3,4,5,0,4,0}
	 moveZeroes(arr)
	 for _ , v := range arr{
	 	fmt.Print(v)
	 }
}
func moveZeroes(nums []int)  {
	var j int
	for i:= 0 ;i< len(nums) ;i++ {
		if nums[i]==0 {
			j=i+1
			in ://给循环加标签
			for j<len(nums) {
				if nums[j] != 0 {
				nums[i]=nums[j]
				nums[j]=0
				break in
				}
				j++;
			}

		}
	}
}
