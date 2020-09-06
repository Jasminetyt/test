package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums=[]int{2,4,11,5}
	fmt.Println(canPartition(nums))
	fmt.Println("-------------")
	var nums2=[]int{0,0,1,2,2,2,2,2,2}
	fmt.Println(wangyibeibao(nums2,10))
}
//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
func canPartition(nums []int) bool {
	var sum int
	for i:=0;i<len(nums);i++{
		sum=sum+nums[i]
	}
	if sum%2!=0{
		return false
	}
	temp:=sum/2
	sort.Ints(nums)
	var dp=make([]bool,temp+1)
	if nums[0]<=temp{
		dp[nums[0]]=true
	}
	for i:=1;i<len(nums);i++{
		for j:=temp-nums[i];j>=0;j--{
			if dp[j]==true{
				dp[j+nums[i]]=true
			}
		}
		//fmt.Println(dp)
	}
	if dp[temp]==true{
		return true
	}
	return false
} //每个值要么取，要么不取

func wangyibeibao(v []int,w int) int { //返回次数
    sort.Ints(v)
    var index int
    var i int
    for i=0;i<len(v);i++{
    	if v[i]!=0{
    		break
		}
	}
    index=i
	var dp=make([]int,w+1)
	for i:=index;i<len(v);i++{
		for j:=0;j<=v[i]&&j<w+1;j++{
			if dp[j]!=0&&j+v[i]<w+1{
				dp[j+v[i]]=dp[j+v[i]]+dp[j]
			}
		}
		if v[i]<w+1{
			dp[v[i]]=dp[v[i]]+1
		}
	}
	var sum int
	for i:=0;i<w+1;i++{
		sum=sum+dp[i]
	}
	return sum+1
}