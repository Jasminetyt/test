package main

import (
	"fmt"
)

func main() {
	var nums=[]int{70, 30, 40, 10, 80, 80,20, 90, 100, 75, 60, 45}

	mergeSort1(nums,0,len(nums)-1)
	fmt.Println(nums)
	heapSort(nums)
	quickSort(0,len(nums)-1,nums)
	fmt.Println(nums)
}
func quickSort(start,end int,nums []int)  {
	if start>=end{
		return
	}
	var i,j=start,end
	temp:=nums[i]
	for i<j{
		for i<j&&nums[j]>temp{
			j--
		}
		if i<j{
			nums[i]=nums[j]
			i++
		}
		for i<j&&nums[i]<temp{
			i++
		}
		if i<j{
			nums[j]=nums[i]
			j--
		}
	}
	nums[i]=temp
	quickSort(start,i-1,nums)
	quickSort(i+1,end,nums)
}
//根节点下标从1开始
func heapSort(nums []int)  {
	//创建堆
	for i:=len(nums)/2-1;i>=0;i--{
		createHeap2(nums,i,len(nums)-1)
	}
	//堆排序
	for i:=len(nums)-1;i>=0;i--{
		nums[i],nums[0]=nums[0],nums[i]
		createHeap2(nums,0,i-1)
	}
	fmt.Println(nums)
}
func createHeap2(nums []int,index,n int)  {
	temp:=nums[index]
	for i:=2*index+1;i<=n;i=i*2+1{
		if i<n&&nums[i]<nums[i+1]{
			i++
		}
		if nums[i]>temp{
			nums[index]=nums[i]
			index=i
		}
	}
	nums[index]=temp
}

func mergeSort1(nums []int,start,end int)  {
	if start<end{
		mid:=(start+end)/2
		mergeSort1(nums,start,mid)
		mergeSort1(nums,mid+1,end)
		merge1(nums,start,mid,end)
	}
}
func merge1(nums []int,start,mid,end int)  {
	var temp=make([]int,end-start+1)
	var index1,index2,i=start,mid+1,0
	for ;index1<=mid&&index2<=end;i++{
		if nums[index1]>nums[index2]{
			temp[i]=nums[index2]
			index2++
		}else {
			temp[i]=nums[index1]
			index1++
		}
	}
	for index1<=mid{
		temp[i]=nums[index1]
		i++
		index1++
	}
	for index2<=end{
		temp[i]=nums[index2]
		index2++
		i++
	}
	for i:=0;i<len(temp);i++{
		nums[start+i]=temp[i]
	}
}
















