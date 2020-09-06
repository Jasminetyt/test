package main

import (
	"fmt"
)

func main() {
	var(
		n,m,temp  int
	)
	fmt.Scanf("%d,%d",&m,&n) //m为行数，n为列数
	//fmt.Printf("m=%d,n=%d\n",m,n)
	num := make([][]int ,m)
	for i:=0;i<m ;i++  {
		for j:=0;j<n ;j++  {
			fmt.Scan(&temp)
			num[i]=append(num[i],temp)
		}
	}
	printMatrixInCircle(num)
	fmt.Println(" ")
	printMatrixInCircleWisely(num)
}
func printMatrixInCircle(num [][]int)  {
	var flag1=len(num[0]) //共多少列
	var flag2=len(num)-1 //共多少行
	var i,j = 0,-1
	for flag1>0 || flag2>0 {
		if flag1>0&&j<len(num[0])-1{
			i,j=printRightLine(i,j+1,flag1,num)
		}
		flag1--
		if flag2>0&&i<len(num)-1 {
			i,j=printDownLine(i+1,j,flag2,num)
		}
		flag2--
		if flag1>0&&j>0&&flag2>=0 {
			i,j=printLeftLine(i,j-1,flag1,num)
		}
		flag1--
		if flag2>0&&i>0&&flag1>=0 {
			i,j=printUpLine(i-1,j,flag2,num)
		}
		flag2--
		if flag1<=0 || flag2<=0{
			return
		}
	}
}
func printRightLine(i,j,flag int,num [][]int) (int, int) {
	//fmt.Println("printRightLine")
	var count=0
	for count < flag  {
		fmt.Printf("%d ",num[i][j+count])
		count++
	}
	return i,j+count-1
}
func printDownLine(i,j,flag int,num [][]int) (int,int) {
	//fmt.Println("printDownLine")
	var count  =0
	for count<flag{
		fmt.Printf("%d ",num[i+count][j])
		count++
	}
	return i+count-1,j
}
func printLeftLine(i,j,flag int,num [][]int) (int,int) {
	//fmt.Println("printLeftLine")
	var count  =0
	for count<flag{
		fmt.Printf("%d ",num[i][j-count])
		count++
	}
	return i,j-count+1
}
func printUpLine(i,j,flag int,num [][]int) (int,int) {
	//fmt.Println("printUpLine")
	var count  =0
	for count<flag{
		fmt.Printf("%d ",num[i-count][j])
		count++
	}
	return i-count+1,j
}

func printMatrixInCircleWisely(num [][]int){
	var start int
	var row,col=len(num),len(num[0])
	for row>2*start&&col>2*start{
		printMatrixInCircle2(num,start,row,col)
		start++
	}
}
func printMatrixInCircle2(num [][]int,start,row,col int){
	endx:=col-1-start
	endy:=row-1-start
	for i:=start;i<=endx;i++{
		fmt.Printf("%d ",num[start][i])
	}
	if endy>start{
		for i:=start+1;i<=endy;i++{
			fmt.Printf("%d ",num[i][endx])
		}
	}
	if endy>start&&endx>start{
		for i:=endx-1;i>=start;i--{
			fmt.Printf("%d ",num[endy][i])
		}
	}
	if endx>start&&start+1<endy{
		for i:=endy-1;i>=start+1;i--{
			fmt.Printf("%d ",num[i][start])
		}
	}
}
