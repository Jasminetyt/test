package main

import (
	"fmt"
	"sort"
)
func main(){
	var num=make([]int,5)
	for i:=0;i<5;i++{
		fmt.Scan(&num[i])
	}
	sort.Ints(num)
	var nums=make([]int,10)
	var count=0
	for i:=0;i<5;i++{
		for j:=i+1;j<5;j++{
			nums[count]=minbeishu(num[i],num[j])
			count++
		}
	}
	sort.Ints(nums)
	var temp=minbeishu(minbeishu(num[0],num[1]),num[2])
	//fmt.Println(nums,temp)
	for i:=0;i<10;i++{
		count=0
		if nums[i]<temp{
			for j:=0;j<5;j++{
				if nums[i]%num[j]==0{
					count++
				}
			}
			if count>=3{
				fmt.Println(nums[i])
				return
			}
		}else{
			fmt.Println(temp)
			return
		}
	}
	fmt.Println(temp)
}

func minbeishu(x,y int) int { //两个数的最小公倍数
	max,step:=y,y
	for {
		if max%x==0&&max%y==0{
			return max
		}
		max=max+step
	}
}
