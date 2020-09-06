package main

import (
	"fmt"
	"math/rand"
	sort2 "sort"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	var nums []int
	for i:=0;i<5;i++{
		num :=rand.Intn(13)
		nums=append(nums,num)
	}
	fmt.Println(nums)
	isContinous(nums)
}
func isContinous(nums []int)  {
	if len(nums)==0 || nums==nil {
		return
	}
	sort2.Ints(nums)
	var numberOfZero int
	for i:=0;i<len(nums)&&nums[i]==0;i++{
			numberOfZero++
	}
	small:=numberOfZero
	big:=small+1
	var numberOfGap int
	for big<len(nums){
		if nums[big]==nums[small] {
			fmt.Println("有对子，不是顺子")
			return
		}
		numberOfGap=numberOfGap+nums[big]-nums[small]-1
		small=big
		big++
	}
	fmt.Println("numberOfZero",numberOfZero,"numberOfGap",numberOfGap)
	if numberOfZero>=numberOfGap {
		fmt.Println("顺子")
	}else {
		fmt.Println("不是顺子")
	}
}
