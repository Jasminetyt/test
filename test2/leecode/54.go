package main

import "fmt"

func main() {
	var matrix=[][]int{{1},{2},{3}}
	fmt.Println(spiralOrder(matrix))
}
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix)==0{
		return nil
	}
	var (
		length1=len(matrix[0])
		length2=len(matrix)
		length3=0
		length4=1
		nums []int
		x int
		y=-1
	)
	for length1>=length3&&length2>=length4{
		y=y+1
		if y>=length1 {
			break
		}
		for;y<length1;y++{ //向右
			nums=append(nums,matrix[x][y])
		}
		y--
		length1--

		x=x+1
		if x>=length2 {
			break
		}
		for ;x<length2;x++{ //向下
			nums=append(nums,matrix[x][y])
		}
		x--
		length2--

		y=y-1
		if y<length3 {
			break
		}
		for ;y>=length3;y--{//向左
			nums=append(nums,matrix[x][y])
		}
		y++
		length3++

		x=x-1
		if x<length4 {
			break
		}
		for ;x>=length4;x--{//向上
			nums=append(nums,matrix[x][y])
		}
		x++
		length4++
	}
	return nums
}