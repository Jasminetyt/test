package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//不改变数组
	//只适用于找出任意一个，不适用找出全部重复数字
	scan := bufio.NewScanner(os.Stdin)
	var count = 0
	var str string
	for scan.Scan(){
		str=scan.Text()
		count++
		if count==1{
			break
		}
	}
	strs:=strings.Split(str," ")
	var num []int
	var temp int
	for i:=0;i<len(strs);i++{
		temp,_=strconv.Atoi(strs[i])
		num=append(num,temp)
	}

	start,end := 1,len(strs)-1
	for start<end{
		mid := (start+end)/2
		count:=countRange(num,start,mid)
		if end==start{
			if count>1{
				fmt.Println(start)
				break
			}else {
				break
			}
		}
		if count>(mid-start+1) {
			end=mid
		}else {
			start=mid+1
		}

	}
}
func countRange(num []int,start,end  int) (count int) {
	count=0
	for i:=0;i<len(num);i++{
		if num[i]>=start&&num[i]<=end{
			count++
		}
	}
	return count
}
