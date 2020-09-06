package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)
type compareNandM2 []string

func (c compareNandM2) Len() int {
	return len(c)
} //长度
func (c compareNandM2) Less(i,j int) bool { //比较方法
	var temp1=c[i]+c[j]
	var temp2=c[j]+c[i]
	result:=strings.Compare(temp1,temp2) //小于为-1，相等为0，大于为1
	if result<=0 {
		return true
	}else {
		return false
	}
}
func (c compareNandM2) Swap(i,j int)  { //交换方法
	c[i],c[j]=c[j],c[i]
}
func main() {
	var strs compareNandM2
	scan :=bufio.NewScanner(os.Stdin)
	var str string
	for scan.Scan(){
		str=scan.Text()
		break
	}
	strs=strings.Split(str,",")
	//var nums=make([]int,len(strs))
	//for i:=0;i<len(nums);i++{
	//	temp,_:=strconv.Atoi(strs[i])
	//	nums[i]=temp
	//}
	sort.Sort(strs)
	for _,value := range strs{
		fmt.Printf("%s",value)
	}
}
