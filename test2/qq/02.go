package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	temp:=int(math.Pow(2,float64(n)))
	var nums=make([]int,temp)
	for i:=0;i<temp;i++{
		fmt.Scan(&nums[i])
	}
	var m int
	fmt.Scan(&m)
	var num2=make([]int,m)
	for i:=0;i<m;i++{
		fmt.Scan(&num2[i])
	}
	var result=make([]int,m)
	for i:=0;i<m;i++{
		resver(nums,num2[i])
		result[i]=getres(nums)
	}
	for i:=0;i<m;i++{
		fmt.Println(result[i])
	}
	//fmt.Println(result)
}
func resver(nums []int,n int)  {
	var temp=int(math.Pow(2,float64(n)))
	for i:=0;i<len(nums);{
		for j:=0;j<temp/2;j++{
			nums[i+j],nums[i+temp-1-j]=nums[i+temp-1-j],nums[i+j]
		}
		i=i+temp
	}
	//fmt.Println(nums)
}
func getres(nums []int)int  {
	var count int
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[j]<nums[i]{
				count++
			}
		}
	}
	return count
}