package main

import "fmt"
//找缺失的第一个正整数
func main() {
	var nums=[]int{2,1,0,-1}//{1}{1,2,3},{1,3,3},{2}
	fmt.Println(firstMissingPositive(nums))
}
func firstMissingPositive(nums []int) int {
	if len(nums)==0 || nums==nil{
		return 1
	}
	var i int
	for i=0;i<len(nums);i++{ //如果当前数组中不存在1，怎返回结果1
		if nums[i]==1{
			break
		}
	}
	if i==len(nums){
		return 1
	}
	length := len(nums)
	for i=0;i<length;i++{ //将数组中小于1，大于数组长度的数都换成1
		if nums[i]<=0 || nums[i]>length{
			nums[i]=1
		}
	}
	for i=0;i<length;{//尽量的将将数组中的数换到换到下标与值相等的位置
		temp := nums[i]-1
		if nums[temp]==temp+1{
			i++
		}else{
			nums[temp],nums[i]=nums[i],nums[temp]
		}
	}

	for i=0;i<length;i++{ //如下标与值不想等，则该下标+1为缺失的最小正整数
		if nums[i]!=i+1{
			return i+1
		}
	}

	return i+1 //{1，2，3}没有数缺失的情况

}