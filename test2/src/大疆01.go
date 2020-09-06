package main

import (
	"fmt"
	"io"
)
func main() {
	var N,A,X int//bug数，效率,最多次数
	var result []int
	for{
		_,err:=fmt.Scan(&N,&A,&X)
		if err==io.EOF {
			break
		}else {
			var num=make([]int,N)
			var sum=0
			for i:=0;i<N;i++{
				fmt.Scan(&num[i])
				sum=sum+num[i]
			}
			var min,temp int
			var index=60*A*X
			if sum<=index{
				//fmt.Println("sum=",sum,"A=",A,sum/A)
				if sum%A==0 {
					result=append(result,sum/A)
				}else {
					result=append(result,sum/A+1)
				}
			}else {
				temp=sum-index
				min=temp+X*60
				if min>8*60 {
					result=append(result,0)
				}else {
					result=append(result,min)
				}
			}

		}
	}
	for i:=0;i<len(result);i++{
		fmt.Println(result[i])
	}
}
