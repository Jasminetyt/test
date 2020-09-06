package main

import "fmt"

func main() {
	//拆成1，2，4，8... (13=1+2+4+6)
	//拆分成0-1背包问题
	var n,W int //数量和重量
	fmt.Scan(&n,&W)
	var v=make([]int,n)//价值
	var num=make([]int,n)//数量
	var w=make([]int,n)//重量
	//转换为0-1背包
	var c int
	var w2,v2 []int
	for i:=0;i<=n;i++{
		c=1
		for num[i]-c>0{
			num[i]=num[i]-c
			w2=append(w2,c*w[i])
			v2=append(v2,c*v[i])
			c=c*2
		}
		w2=append(w2,num[i]*w[i])
		v2=append(v2,num[i]*v[i])
	}
	//0-1背包
}
