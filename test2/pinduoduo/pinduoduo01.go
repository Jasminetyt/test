package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s =make([]string,2)
	var count int
	e1 := bufio.NewScanner(os.Stdin)
	for e1.Scan()  {
		s[count]=e1.Text()
		count++
		if count==2 {
			break
		}
	}
	s11:=strings.Split(s[0]," ")
	s22 := strings.Split(s[1]," ")
    var A,B []int
	for i:=0;i<len(s11);i++{
		temp,_:=strconv.Atoi(s11[i])
		A=append(A,temp)
	}
	for i:=0;i<len(s22);i++{
		temp,_:=strconv.Atoi(s22[i])
		B=append(B,temp)
	}
	if len(A)<1 || A==nil{
		fmt.Println("NO")
		return
	}
	max:=B[0]
	if len(A)==1{
		for i:=0;i<len(B);i++{
			if B[i]>max {
				max=B[i]
			}
		}
		fmt.Println(max)
		return
	}

	var first int
	var index int
	var last int
	for i:=1;i<len(A);i++{
		if A[i-1]>=A[i]{
			index=i
			break
		}
	}
	//fmt.Println("index=",index)
	if index+1<len(A){
		if A[index+1]<=A[index-1]{
			index=index-1
		}
		if index-1>=0{
			first=A[index-1]
		}//index==0
		last=A[index+1]
	}else {
		first=A[index-1]
		last=A[index]
	}
	//fmt.Println(first,last)
	var temp2=first
	for i:=0;i<len(B);i++{
		if ((B[i]>first|| first==0) &&B[i]<last) || (last==A[index]&&(B[i]>first || first==0)){
			if B[i]>temp2 {
				temp2=B[i]
			}
		}
	}
	if temp2==first{
		fmt.Println("NO")
	}else {
		A[index]=temp2
		fmt.Println(A)
	}
}
