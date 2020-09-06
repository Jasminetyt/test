package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan:=bufio.NewScanner(os.Stdin)
	var str string
	for scan.Scan(){
		str=scan.Text()
		break
	}
	strs:=strings.Split(str," ")
	var num1=make([]int,len(strs))
	for i:=0;i<len(strs);i++{
		temp,_:=strconv.Atoi(strs[i])
		num1[i]=temp
	}
	for scan.Scan(){
		str=scan.Text()
		break
	}
	strs=strings.Split(str," ")
	var num2=make([]int,len(strs))
	for i:=0;i<len(strs);i++{
		temp,_:=strconv.Atoi(strs[i])
		num2[i]=temp
	}
	var mm [][]int
	mm=append(mm,[]int{num1[0],num2[0]})
	for i:=1;i<len(num1);i++{
		if ! comm(mm,num1[i],num2[i]){
			mm=append(mm,[]int{num1[i],num2[i]})
		}
	}
//	fmt.Println(mm)
	fmt.Println(len(mm))
}
func comm(mm [][]int,x,y int) bool  {
	for i:=0;i<len(mm);i++{
		if x>mm[i][1]{
			mm[i][0]=x
			mm[i][1]=y
			return true
		}
	}
	return false
}
