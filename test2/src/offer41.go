package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var count=0
	var mid int
	var max,min []int
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	for scan.Scan(){
		if scan.Text()=="end" {
			break;
		}
		num,_:=strconv.Atoi(scan.Text())
		count++
		if count%2==0 {
			if len(max)==0 {
				min=append(min,num)
			}else {
				temp := maxHeap(max)
				if temp>num {
					min=append(min,temp)
					max[0]=num
				}else {
					min=append(min,num)
				}
			}
		}else {
			if len(min)==0 {
				max=append(max,num)
			}else {
				temp:=minHeap(min)
				if num>temp {
					max=append(max,temp)
					min[0]=num
				}else {
					max=append(max,num)
				}
			}
		}
		if count%2==0 {
			mid = (maxHeap(max)+minHeap(min))/2
		}else {
			mid=maxHeap(max)
		}
		fmt.Println("中位数为：",mid)
	}
}
func initheapode(length int) (int,int, int) {
	var left, right,parent int
	if length%2==0 {
		right = length
		left = right-1
	}else {
		right = -1
		left =length
	}
	parent = left/2
	return left,right,parent
}
func maxHeap(nums []int) int {
	length := len(nums)-1
	left ,right,parent := initheapode(length)
	for left>0{
		if nums[left]>nums[parent] {
			nums[parent],nums[left]=nums[left],nums[parent]
		}
		if right!=-1&&nums[right]>nums[parent] {
			nums[parent],nums[right]=nums[right],nums[parent]
		}
		right=left-1
		left=right-1
		parent=left/2
	}
	return nums[0]
}
func minHeap(nums []int) int {
	length := len(nums)-1
	left,right,parent := initheapode(length)
	for left>0{
		if nums[left]<nums[parent] {
			nums[parent],nums[left]=nums[left],nums[parent]
		}
		if right!=-1&&nums[right]<nums[parent] {
			nums[parent],nums[right]=nums[right],nums[parent]
		}
		right=left-1
		left=right-1
		parent=left/2
	}
	return nums[0]
}