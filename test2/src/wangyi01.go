package main

import "fmt"

func main() {
	var t int //t组数据
	fmt.Scan(&t)
	var n,m,k int //广场大小和障碍物个数
	var result=make([]string,t)
	for i:=0;i<t;i++{
		fmt.Scan(&n,&m,&k)
		result[i]=input(n,m,k)
	}
	for i:=0;i<t;i++{
		fmt.Println(result[i])
	}
}
func input(n,m,k int) string {
	var x,y int
	var temp=make([][]int,n)
	for i:=0;i<n;i++{
		temp[i]=make([]int,m)
	}
	for i:=0;i<k;i++{
		fmt.Scan(&x,&y)
		temp[x-1][y-1]=1
	}
	fmt.Scan(&x,&y)
	//fmt.Println(temp,x,y)
	for i:=0;i+x<=n;i++{
		for j:=0;j+y<=m;j++{
			if check(i,j,x,y,temp){
				return "YES"
			}

		}
	}
	return "NO"
}
func check(startx,starty,x,y int,temp [][]int) bool {
	//fmt.Println("startx=",startx,"starty=",starty)
	var i,j int
	//var count1,count2 int
	for i=startx;i<len(temp)&&i<startx+x;i++{
		for j=starty;j<len(temp[0])&&j<starty+y;j++{
			if temp[i][j]==1{
				//fmt.Println(i,j,"==false")
				//count2=0
				return false
			}
		}
	}
	if i==x+startx&&j==y+starty{
		return true
	}
	if i>=len(temp) || j>=len(temp[0]){
		//fmt.Println("...")
		return false
	}
	return true
}
