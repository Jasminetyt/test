package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var nums []int
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	for scan.Scan(){
		if scan.Text()=="end" {
			break;
		}
		temp ,_ :=strconv.Atoi(scan.Text())
		nums=append(nums,temp)
	}
	if len(nums)<=0 {
		return
	}
	getLeastNumber1(5,nums)
	fmt.Println(" ")
	getLeastNumber2(5,nums)
	fmt.Println("")
	getLeastNumber3(5,nums)
	getLeastNumber4(5,nums)
}
func getLeastNumber1(k int,nums []int)  {
	start ,end := 0,len(nums)-1
	index := partition2(start,end,nums)
	for index+1!=k{
		if index+1<k{
			start=index+1
		}
		if index+1>k {
			end=index-1
		}
		index=partition2(start,end,nums)
	}
	for i:=0;i<=index ;i++  {
		fmt.Printf("%d ",nums[i])
	}
}
func partition2(start,end int ,nums []int) int {
	temp := start
	start=start+1
	for start<end{
		for nums[start]<=nums[temp]{
			start++
		}
		for nums[end]>nums[temp]{
			end--
		}
		if start<end {
			nums[start],nums[end]=nums[end],nums[start]
			start++
			end--
		}
	}
	nums[temp],nums[end]=nums[end],nums[temp]
	return end
}
func getLeastNumber2(k int,nums []int)  {
	heap :=createMaxHeap(nums[:k])
	for i:=k;i<len(nums);i++{
		if nums[i]<heap[0] {
			heap[0],nums[i]=nums[i],heap[0]
			heap=createMaxHeap(heap)
		}
	}
	for i:=0;i<k ;i++  {
		fmt.Printf("%d ",heap[i])
	}
}
func createMaxHeap(nums []int) []int {
	var right,left,parent int
	if (len(nums)-1)%2==0 {
		right=len(nums)-1
		left=right-1
		parent=left/2
	}else {
		right=-1
		left=len(nums)-1
		parent=left/2
	}
	for left>0{
		if nums[left]>nums[parent] {
			if right!=-1&&nums[left]>nums[right] || right==-1{
				nums[left],nums[parent]=nums[parent],nums[left]
			}
		}
		if right!=-1&&nums[right]>nums[parent]&&nums[right]>nums[left] {
			nums[right],nums[parent]=nums[parent],nums[right]
		}
		right=left-1
		left=right-1
		parent=left/2
	}
	return nums
}

func getLeastNumber3(k int,nums []int){
	if len(nums)<k{
		fmt.Println("no exist")
	}
	var start,end=0,len(nums)-1
	var index int
	for start<end{
		index=partition3(start,end,nums)
		if index==k{
			break
		}
		if index<k{
			start=index+1
		}
		if index>k{
			end=index-1
		}
	}
fmt.Println(nums[:k])
}
func partition3(start,end int ,nums []int) int {
	index:=nums[start]
	for start<end{
		for start<end&&nums[end]>index{
			end--
		}
		if start<end{
			nums[start]=nums[end]
			start++
		}
		for start<end&&nums[start]<index{
			start++
		}
		if start<end{
			nums[end]=nums[start]
			end--
		}
	}
	nums[start]=index
	return start
}

func getLeastNumber4(k int,nums []int)  {
	createMaxHeap2(nums,k)
	for i:=k;i<len(nums);i++{
		if nums[i]<nums[0]{
			nums[0],nums[k-1]=nums[k-1],nums[0]
			nums[k-1],nums[i]=nums[i],nums[k-1]
			createMaxHeap2(nums,k)
		}
	}
	fmt.Println(nums[0:k])
}
func createMaxHeap2(nums []int,k int)  {
	var parent,left,right int
		right=k-1
		left=right-1
		parent=left/2
	for left>0{
		if nums[left]>nums[right]&&nums[left]>nums[parent]{
			nums[parent],nums[left]=nums[left],nums[parent]
	    }
		if nums[right]>=nums[left]&&nums[right]>nums[parent]{
			nums[parent],nums[right]=nums[right],nums[parent]
		}
		right=left-1
		left=right-1
		parent=left/2
	}
}