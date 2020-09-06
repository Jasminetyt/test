package main

import "fmt"

func main() {
	var n=5
	fmt.Println(solveNQueens(n))
}
func solveNQueens(n int) [][]string {
	var index=0
	var result [][]string
	var temp=make([][]int,n)
	for i:=0;i<n;i++{
		temp[i]=make([]int,n)
	}
	solveNQueens1(index,n,&result,temp)
	return result
}

func solveNQueens1(index,n int,result *[][]string,temp [][]int)  {
	if index==n{
		getresult(result,temp)
		return
	}
	for i:=0;i<n;i++{
		if temp[index][i]==0 {
			copyvalue(index,i,1,-(index+1),temp)
			solveNQueens1(index+1,n,result,temp)
			deletevalue(index,i,0,-(index+1),temp)
		}

	}
}
func copyvalue(index ,i,queen,no int,temp [][]int)  {
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
func deletevalue(index ,i ,queen,no int,temp [][]int)  {
	temp[index][i]=queen
	for x:=0;x<len(temp);x++{
		for y:=0;y<len(temp);y++{
			if temp[x][y]==no {
				temp[x][y]=0
			}
		}
	}
}
func getresult(result *[][]string,temp [][]int)  {
	var temp2=make([]string,len(temp))
	for i:=0;i<len(temp);i++{
		var str string
		for j:=0;j<len(temp);j++{
			if temp[i][j]!=1 {
				str=str+"."
			}else {
				str=str+"Q"
			}
		}
		temp2[i]=str
	}
	*result=append(*result,temp2)
}
