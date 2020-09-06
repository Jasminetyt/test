package main

import (
	"bufio"
	"fmt"
	"os"
	s1 "sort"
	"strconv"
	"strings"
)

type compareNandM []string

func (c compareNandM) Len() int {
	return len(c)
} //长度
func (c compareNandM) Less(i,j int) bool { //比较方法
	var temp1=c[i]+c[j]
	var temp2=c[j]+c[i]
	result:=strings.Compare(temp1,temp2) //小于为-1，相等为0，大于为1
	if result<=0 {
		return true
	}else {
		return false
	}
}
func (c compareNandM) Swap(i,j int)  { //交换方法
	c[i],c[j]=c[j],c[i]
}
func main() {
	var nums compareNandM
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	for scan.Scan(){
		if scan.Text()=="end"{
			break;
		}
		nums=append(nums,scan.Text())
	}
	if len(nums)<=0 {
		return
	}
	s1.Sort(nums)
	for _,value := range nums{
		fmt.Printf("%s",value)
	}
	fmt.Println("")
	getlittilenum()
}
type comaprenum []string

func (s comaprenum) Less(i,j int) bool {
	temp1:=s[i]+s[j]
	temp2:=s[j]+s[i]
	if strings.Compare(temp1,temp2)<0{
		return true
	}
	return false
}
func (s comaprenum) Len() int {
	return len(s)
}
func (s comaprenum) Swap(i,j int)  {
	s[i],s[j]=s[j],s[i]
}

func getlittilenum(){
	var num comaprenum
	var num2=[]int{3,32,321}
	for i:=0;i<len(num2);i++{
		temp:=strconv.Itoa(num2[i])
		num=append(num,temp)
	}
	s1.Sort(num)
	fmt.Println(num)
}
