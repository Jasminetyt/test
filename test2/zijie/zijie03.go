package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n,m,q int //n代表方块个数，m代表数字最大不超过m，q代表次数
	fmt.Scan(&n,&m,&q)
	var strs=make([]string,n)
	var counts=make([][]int,q)
	for i:=0;i<len(strs);i++{
		fmt.Scan(&strs[i])
	}
	for i:=0;i<len(counts);i++{
		counts[i]=make([]int,2)
	}
	for i:=0;i<len(counts);i++{
		for j:=0;j<2;j++{
			fmt.Scan(&counts[i][j])
		}
	}
	var result=make([]int,q)
	for i:=0;i<len(counts);i++{
		getcount(counts[i][0],counts[i][1],strs,i,result)
	}
	for i:=0;i<len(result);i++{
		fmt.Println(result[i])
	}
}
func getcount(start,end int,strs []string,index int,result []int){
	var i=start-1
	var temp1=strs[i:end]
	var count int
	var flag bool
	for{
		if temp1[i]==">"{
			i++
			flag=false
			if temp1[i]==">" || temp1[i]=="<"{
				delete(i,temp1)
			}
		}
		if temp1[i]=="<"{
			i--
			flag=true
			if temp1[i]==">" || temp1[i]=="<"{
				delete(i,temp1)
			}
		}
		if i<start-1 || i>end-1 {
			break
		}
		value,_:=strconv.Atoi(temp1[i])
		count=count+value
		value=value-1
		if value==0 {
			delete(i,temp1)
		}else {
			temp1[i]=strconv.Itoa(value)
		}
		if !flag{
			i++
		}else {
				i--
		}

		if i==start-1 || i==end-1{
			if temp1[i]!="<"&&temp1[i]!=">"{
				value,_:=strconv.Atoi(temp1[i])
				count=count+value
				break
			}

		}
		if len(temp1)<=1{
			break;
		}

	}
	result[i]=count
}
func delete(i int,temp1 []string)  {
	temp2:=temp1[i+1:]
	temp3:=temp1[0:i]
	temp1=make([]string,len(temp2)+len(temp3))
	var j int
	for j=0;j<len(temp3);j++{
		temp1[j]=temp3[j]
	}
	for z:=0;z<len(temp2);z,j=z+1,j+1{
		temp1[j]=temp2[z]
	}
}
