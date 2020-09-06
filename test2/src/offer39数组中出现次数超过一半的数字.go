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
			break
		}
		temp ,_ := strconv.Atoi(scan.Text())
		nums=append(nums,temp)
	}
	if len(nums)<=0 {
		fmt.Println("输入的数组有误")
	}
	flag , value := moreThanHalfNum1(nums)
	if flag {
		fmt.Println(value)
	}
	flag , value = moreThanHalfNum2(nums)
	if flag {
		fmt.Println(value)
	}
	moreThanHalfNum3(nums)
}
//排序后，在数组中出现次数超过一半的数字一定是第n/2大的数字
//获得第n/2大的数字后，校验该数字出现次数是否超过数组的一半
func moreThanHalfNum1(nums []int) (bool, int ){
	index,start , end := 0,0,len(nums)-1
	for index!=(len(nums)/2)&&start<end{
		index=partition(nums,start,end)
		if index==(len(nums)/2) {
			break
		}
		if index<(len(nums)/2){
			start=index+1
		}
		if index>(len(nums)/2) {
			end=index-1
		}
	}
	//校验
	count :=0
	for i:=0;i<len(nums);i++{
		if nums[i]==nums[index] {
			count++
		}
	}
	if count>(len(nums)/2) {
		return true,nums[index]
	}else {
		return false,0
	}
}
func partition(nums []int,start,end int) int {
	temp := start
	for start<end{
		for start<end&&nums[start]<=nums[temp]{
			start++
		}
		for start<end&& nums[end]>nums[temp]{
			end--
		}
		if start<end{
			nums[end],nums[start]=nums[start],nums[end]
			start++
			end--
		}
	}
	end--
	nums[temp],nums[end]=nums[end],nums[temp]
	return end
}

func moreThanHalfNum2(nums []int)(bool,int){
	var temp,count =nums[0],0
	for i:=0;i<len(nums);i++{
		if count==0 {
			temp=nums[i]
			//fmt.Println(temp)
			count=1
		}else if nums[i]==temp {
			count++
		}else{
			count--
		}
	}
	c:=0
	if count>=1 {
		//还需校验
		for i:=0;i<len(nums);i++{
			if nums[i]==temp {
				c++
			}
		}
		if c*2>len(nums) {
			return true	,temp
		}else {
			return false,0
		}

	}else {
		return false ,0
	}

}

func moreThanHalfNum3(nums []int)  {
	var temp,count =nums[0],1
	for i:=1;i<len(nums);i++{
		if nums[i]==temp{
			count++
		}else {
				count--
				if count==0{
					temp=nums[i]
					count=1
				}
		}
	}
	//还需校验，如2 2 1 3
	count=0
	for i:=0;i<len(nums);i++{
		if nums[i]==temp{
			count++
		}
	}
	if count*2>len(nums){
		fmt.Println(temp)
	}else {
		fmt.Println("no exsist")
	}

}