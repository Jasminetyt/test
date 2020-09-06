package main

import "fmt"

func main() {
	fmt.Println(totalNQueens(5))
}
func totalNQueens(n int) int {
	var count int
	var temp=make([][]int,n)
	for i:=0;i<n;i++{
		temp[i]=make([]int,n)
	}
	solveNQueens2(0,n,temp,&count)
	return count
}

func solveNQueens2(index,n int,temp [][]int,count *int){
	if index==n{
		*count++
		return
	}
	for i:=0;i<n;i++{
		if temp[index][i]==0 {
			copyvalue2(index,i,1,-(index+1),temp)
			solveNQueens2(index+1,n,temp,count)
			deletevalue2(index,i,0,-(index+1),temp)
		}
	}
}
func copyvalue2(index ,i,queen,no int,temp [][]int)  {
	temp[index][i]=queen
	for j:=0;j<len(temp);j++{
		if j!=i&&temp[index][j]==0 {
			temp[index][j]=no
		}
	}
	for j:=0;j<len(temp);j++{
		if j!=index&&temp[j][i]==0 {
			temp[j][i]=no
		}
	}

	var x,y int
	if index>i {
		x=index-i
		y=0
	}else {
		x=0
		y=i-index
	}
	for;x<len(temp)&&y<len(temp);{
		if x!=index&&y!=i&&temp[x][y]==0 {
			temp[x][y]=no
		}
		x++
		y++
	}
	if index<len(temp)-1-i {
		x=0
		y=i+index
	}else {
		y=len(temp)-1
		x=index-len(temp)+1+i
	}
	for ;x<len(temp)&&y>=0;{
		if x!=index&&y!=i&&temp[x][y]==0{
			temp[x][y]=no
		}
		x++
		y--
	}

}
func deletevalue2(index ,i ,queen,no int,temp [][]int)  {
	temp[index][i]=queen
	for x:=0;x<len(temp);x++{
		for y:=0;y<len(temp);y++{
			if temp[x][y]==no {
				temp[x][y]=0
			}
		}
	}
}