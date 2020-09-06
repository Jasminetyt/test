package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums=[]int{1, 0, -1, 0, -2, 2}
	var target=0
	fmt.Println(fourSum(nums,target))
}
func fourSum(nums []int, target int) [][]int {
	if len(nums)<4 || nums==nil{
		return nil
	}
	var result =make([][]int,0)
	var temp=make([]int,4)
	sort.Ints(nums)
	getsum(&result,target,0,0,temp,nums)
	return result
}
func getsum(result *[][]int,sum,count,k int,temp,num []int)  {
	if count==4&&sum==0 {
		var temp2=make([]int,4) //新构造一个数组，否则将指向temp
		for i:=0;i<len(temp);i++{
			temp2[i]=temp[i]
		}
		*result=append(*result,temp2)
	}
	if count<4 {
		for i:=k;i<len(num);{
			sum=sum-num[i]
			temp[count]=num[i]
			count++
			getsum(result,sum,count,i+1,temp,num)
			sum=sum+num[i]
			for i+1<len(num)&&num[i]==num[i+1]{
				i++
			}
			i++
			count--

		}
	}
}
