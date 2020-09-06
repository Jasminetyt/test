package main

import (
	"fmt"

)

func main() {
	var nums=make([][]int,5)
	for i:=0;i<5;i++{
		nums[i]=make([]int,5)
	}
		for i:=0;i<5;i++{
			for j:=0;j<5;j++{
				fmt.Scan(&nums[i][j])
			}
		}

	var flag=true
	for flag{
		flag=false
		for i:=0;i<5;i++{
			for j:=0;j<5;j++{
				if nums[i][j]!=0{
					if del(nums,i,j){
						flag=true
						down(nums)
					}
				}
			//	fmt.Println("down",nums)
			}
		}
	}
	var count int
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if nums[i][j]!=0{
				count++
			}
		}
	}
	fmt.Println(count)

}
func del(nums [][]int,i,j int) bool {
	temp:=nums[i][j]
	var count =1
	if i-1>=0&&nums[i-1][j]==temp{ //left
		count++
	}
	if i+1<5&&nums[i+1][j]==temp{ //right
		count++
	}
	if j-1>=0&&nums[i][j-1]==temp{ //up
		count++
	}
	if j+1<5&&nums[i][j+1]==temp{ //down
		count++
	}
	if count>=3{
		nums[i][j]=0
		del2(i,j,temp,nums)
	}
	if count>=3{
		return true
	}else {
		return false
	}
}
func del2(i,j,index int,nums [][]int)  {
	//fmt.Println(nums)
	if i-1>=0&&nums[i-1][j]==index{
		nums[i-1][j]=0
		del2(i-1,j,index,nums)
	}
	if i+1<5&&nums[i+1][j]==index{
		nums[i+1][j]=0
		del2(i+1,j,index,nums)
	}
	if j-1>=0&&nums[i][j-1]==index{
		nums[i][j-1]=0
		del2(i,j-1,index,nums)
	}
	if j+1<5&&nums[i][j+1]==index{
		nums[i][j+1]=0
		del2(i,j+1,index,nums)
	}
}
func down(nums [][]int)  {
	var x int
	for i:=3;i>=0;i--{
		for j:=0;j<5;j++{
			if nums[i][j]!=0{
				temp:=nums[i][j]
				nums[i][j]=0
				for x=i+1;x<5;x++{
					if nums[x][j]!=0{
						break
					}
				}
				if x==i+1{
					nums[i][j]=temp
				}else {
					nums[x-1][j]=temp
				}


			}
		}
	}
}