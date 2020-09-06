package main

import "fmt"

func main() {
	nums := [][]byte{{'a','b','t','g'},{'c','f','c','s'},{'j','d','e','h'}}
	var flag  = make([][]bool,4)
	for i:=0;i<4;i++{
		flag[i]=make([]bool,4)
	}
	str :="bfce"
	strs:=[]byte(str)
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums);j++{
			if nums[i][j]==strs[0]{
				flag[i][j]=true
				fmt.Println(string(strs[0]),i,j)
			if 	findpath(nums,strs,flag,1,i,j){
				fmt.Println(true)
				return
			}
			}
		}
	}
	fmt.Println(false)
}

func findpath(nums [][]byte,strs []byte,flag [][]bool,count,i,j int) bool {
	if count>=len(strs){
		return true
	}
	if i-1>=0&&nums[i-1][j]==strs[count]&&!flag[i-1][j]{//上
		flag[i-1][j]=true
		fmt.Println(string(strs[count]),i-1,j)
		if findpath(nums,strs,flag,count+1,i-1,j){
			return true
		}
	}
	if i+1<len(nums)&&nums[i+1][j]==strs[count]&&!flag[i+1][j]{//下
		flag[i+1][j]=true
		fmt.Println(string(strs[count]),i+1,j)
		if findpath(nums,strs,flag,count+1,i+1,j){
			return true
		}
	}
	if j-1>=0&&nums[i][j-1]==strs[count]&&!flag[i][j-1]{//左
		flag[i][j-1]=true
		fmt.Println(string(strs[count]),i,j-1)
		if findpath(nums,strs,flag,count+1,i,j-1){
			return true
		}
	}
	if j+1<len(nums[0])&&nums[i][j+1]==strs[count]&&!flag[i][j+1]{//右
		flag[i][j+1]=true
		fmt.Println(string(strs[count]),i,j+1)
		if findpath(nums,strs,flag,count+1,i,j+1){
			return true
		}
	}
	fmt.Println("return")
	return false
}
