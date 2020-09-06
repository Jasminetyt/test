package main

import "fmt"

func main() {
	var m,n int //m行n列
	fmt.Scan(&m,&n)
	if m<=0 || n<=0{
		return
	}
	var num=make([][]int,m)
	for i:=0;i<m;i++{
		num[i]=make([]int,n)
		for j:=0;j<n;j++{
			fmt.Scan(&num[i][j])
		}
	}
	fmt.Println("intput search target")
	var target int
	fmt.Scan(&target)

	var row,col=0,n-1
	for row<n&&col>=0{
		if num[row][col]==target{
			fmt.Println("exsist")
			break
		}else if num[row][col]>target {
			col--
		}else {
				row++
		}
	}
	fmt.Println("not exsist")
}
