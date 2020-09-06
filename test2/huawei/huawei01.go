package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var str string
	scan:=bufio.NewScanner(os.Stdin)
	for scan.Scan(){
		str=scan.Text()
		break
	}
	strs:=strings.Split(str," ")
	var nums=make([]int,len(strs))
	for i:=0;i<len(strs);i++{
		temp,_:=strconv.Atoi(strs[i])
		nums[i]=temp
	}
	if len(nums)==2&&nums[1]==0{
		fmt.Println(-1)
		return
	}
	var i=getMax(nums)
	//fmt.Println(i,nums[i])
	count:=1
	for ;i<len(nums);{
		if nums[i]+i+1>=len(nums){
			count++
			fmt.Println(count)
			return
		}
		i=i+nums[i]
		count++
		//fmt.Println(i,nums[i],count)
		if nums[i]==0{
			fmt.Println(-1)
			return
		}
	}

}
func getMax(nums []int) int {
	var max =1
	//fmt.Println("dd",len(nums)/2)
	for i:=1;i<len(nums)/2;i++{
		if nums[i]+i>nums[max]+max{
			max=i
		}
	}
	return max
}
