package main

import (
	"fmt"
	"math"
)
var total=0
func main() {
	var nums=[]int{7,5,6,4}
	if len(nums)<=0 || nums==nil{
		return
	}
	//inversePairs(nums)
	//fmt.Println(total)
	var count=0
	inversePairs2(nums,0,len(nums)-1,&count)
	fmt.Println(count)
}
func inversePairs(nums []int)  {
	mergesort(nums)
}

func mergesort(nums []int)  {
	index := float64(0)
	temp := int(math.Pow(2,index))
	for temp<len(nums){
		for i:=0;i<len(nums);i=i+2*temp{
			ssort(nums,i,i+temp,i+2*temp) //start-mid,mid-end (不包括后者）
		}
		index++
		temp = int(math.Pow(2,index))
	}
}
func ssort(nums []int,start,mid,end int)  {
	if end>len(nums) {
		end=len(nums)
	}
	if mid>len(nums) {
		return
	}
	var temp=make([]int,end-start)
	index := 0
	var i,j int
	for i,j=start,mid;i<mid&&j<end;{
		if nums[i]<=nums[j] {
			temp[index]=nums[i]
			i++
			total=total+j-mid
		}else {
			temp[index]=nums[j]
			j++
		}
		index++
	}
	for i<mid{
		temp[index]=nums[i]
		i++
		index++
		total=total+j-mid
	}
	for j<end{
		temp[index]=nums[j]
		j++
		index++
	}
	for i=0;i<len(temp) ;i++  {
		nums[start]=temp[i]
		start++
	}
}

func inversePairs2(nums []int,start,end int,count *int){
	if start<end{
		var mid=(start+end)/2
		inversePairs2(nums,start,mid,count)
		inversePairs2(nums,mid+1,end,count)
		merge2(nums,start,mid,end,count)
	}
}
func merge2(nums []int,start,mid,end int,count *int)  {
	var temp=make([]int,end-start+1)
	var index1,index2,i=mid,end,end-start
	var length=end-mid
	for ;index1>=start&&index2>mid;{
		if nums[index1]>nums[index2]{
			temp[i]=nums[index1]
			i--
			index1--
			*count=*count+length
		}else {
			temp[i]=nums[index2]
			i--
			index2--
			length--
		}
	}
	for index1>=start{
		temp[i]=nums[index1]
		i--
		index1--
	}
	for index2>mid{
		temp[i]=nums[index2]
		i--
		index2--
	}
	for i:=0;i<len(temp);i++{
		nums[i+start]=temp[i]
	}
}
