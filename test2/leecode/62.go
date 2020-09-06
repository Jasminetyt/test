package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3,3))
}
//func uniquePaths(m int, n int) int {
//	if m==0 || n==0{
//		return 0
//	}
//	var count=0
//	uniquePaths2(m,n,&count,0,0)
//	return count
//}
//func uniquePaths2(m int, n int,count *int,x,y int)  {
//	fmt.Println(x,y)
//	if x==m || y==n{
//		return
//	}
//	if x==m-1&&y==n-1{
//		(*count)++
//		return
//	}
//		uniquePaths2(m,n,count,x+1,y)
//		uniquePaths2(m,n,count,x,y+1)
//}
func uniquePaths(m int, n int) int {
	var col1=make([]int,n)
	var col2=make([]int,n)
	for i:=0;i<n;i++{
		col1[i]=1
		col2[i]=1
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			col2[j]=col2[j-1]+col1[j]
		}
		copy(col1,col2)
	}
	return col2[n-1]
}