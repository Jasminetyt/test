package main

import "fmt"

func main() {
	var M,N int
	fmt.Scan(&M,&N)
	var ik,jk int
	var num=make([][]bool,M+1)
	var dp=make([][]int,M+1)
	for i:=0;i<=M;i++{
		num[i]=make([]bool,N+1)
		dp[i]=make([]int,N+1)
	}
	for {
		n, _ := fmt.Scan(&ik,&jk)
		if n == 0 {
			break
		} else {
			num[ik][jk]=true
		}
	}
	//num[0][2]=true
	//fmt.Println(num)
	//var count int
	//getpathcount(num,&count,0,0)
	//fmt.Println(count)
	dp[0][0]=1
	for i:=1;i<M+1;i++{
		if !num[i][0]{
			dp[i][0]=1
		}
	}
	for i:=1;i<N+1;i++{
		if !num[0][i]{
			dp[0][i]=1
		}
	}
	for i:=1;i<M+1;i++{
		for j:=1;j<N+1;j++{
			if !num[i][j]{
				dp[i][j]=dp[i-1][j]+dp[i][j-1]
			}
		}
	}
	fmt.Println(dp[M][N])


}
//func getpathcount(flag [][]bool,count *int,x,y int)  {
//	if x==len(flag)-1&&y==len(flag[0])-1{
//		*count=*count+1
//		return
//	}
//	if x>=len(flag)|| y>=len(flag[0]){
//		return
//	}
//	if x+1<len(flag)&&!flag[x+1][y]{
//		getpathcount(flag,count,x+1,y)
//	}
//	if y+1<len(flag[0])&&!flag[x][y+1]{
//		getpathcount(flag,count,x,y+1)
//	}
//}