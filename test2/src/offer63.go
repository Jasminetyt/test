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
		if scan.Text()=="end"{
			break
		}
		num ,_ := strconv.Atoi(scan.Text())
		nums=append(nums,num)
	}
	maxDiff(nums)
}
func maxDiff(nums []int)  {
	if len(nums)<2 || nums==nil {
		return
	}
	var min=nums[0]
	var diff =nums[1]-nums[0]
	for i:=2;i<len(nums);i++{
		temp := nums[i]-min
		if temp>diff {
			diff=temp
		}
		if nums[i]<min{
			min=nums[i]
		}
	}
	fmt.Println("differ=",diff)
}
func maxDiff2(nums []int){
	var min=nums[0]
	var max=nums[1]-min
	for i:=2;i<len(nums);i++{
		if nums[i-1]<min{
			min=nums[i-1]
		}
		temp:=nums[i]-min
		if temp>max{
			max=temp
		}
	}
}