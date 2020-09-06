package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n=3
	var k=3
	fmt.Println(getPermutation(n,k))
}
func getPermutation(n int, k int) string {
	var temp1=make([]string,n)
	var flag=make([]bool,n)
	var result []string
	var count int
	var sig=false
	for i:=1;i<=n;i++{
		temp1[i-1]=strconv.Itoa(i)
		flag[i-1]=true
	}
	return getPermutation2(k,temp1,flag,&result,&count,&sig)
}
func getPermutation2(k int,temp1 []string,flag []bool,result *[]string,count *int,sig *bool) string{
	//fmt.Println(temp1,flag)
	if len(*result)==len(temp1){
		*count++
		if *count==k{
			*sig=true
			var str string
			for i:=0;i<len(*result);i++{
				str=str+(*result)[i]
			}
			return str
		}
		return ""
	}
	for i:=0;i<len(temp1);i++{
		if flag[i] {
			*result=append(*result,temp1[i])
			flag[i]=false
			str :=getPermutation2(k ,temp1,flag,result,count,sig)
			if *sig{
				return str
			}
			flag[i]=true
			*result=(*result)[:len(*result)-1]
		}
	}
	return ""
}