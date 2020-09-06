package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)
	strs:=[]byte(str)
	var A []int
	var B []int
	var R int
	var i int
	for i=3;i<len(strs);i++{
		if strs[i]=='}'{
			break
		}
		if strs[i]==','{
			continue
		}
		temp,_:=strconv.Atoi(string(str[i]))
		A=append(A,temp)
	}
	for i=i+5;i<len(strs);i++{
		if strs[i]=='}'{
			break
		}
		if strs[i]==','{
			continue
		}
		temp,_:=strconv.Atoi(string(str[i]))
		B=append(B,temp)
	}
	i=i+4
	if i<len(strs){
		temp,_:=strconv.Atoi(string(str[i]))
		R=temp
	}
	var flag bool
	for i:=0;i<len(A);i++{
		flag=false
		for j:=0;j<len(B);j++{
			if A[i]<=B[j]&&B[j]-A[i]<=R{
				fmt.Printf("(%d,%d)",A[i],B[j])
				flag=true
			}
		}
		if flag==false{
			fmt.Println(A[i],"findmin")
			findmin(B,A[i])
		}

	}
}
func findmin(B []int,index int)  {
	for i:=0;i<len(B);i++{
		if B[i]>=index{
			fmt.Printf("(%d,%d)",index,B[i])
			break
		}
	}

}
