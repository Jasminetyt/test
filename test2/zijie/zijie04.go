package main

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n==0{
		return
	}
	chars:=[]byte(strconv.Itoa(n))
	var list=list.New()
var result []string
 translate2(chars,list,len(chars)-1,&result)
	sort.Strings(result)
	for i:=0;i<len(result);i++{
		fmt.Println(result[i])
	}
}

func translate2(chars []byte,list *list.List,i int,result *[]string){
	if i<0{
		var str string
		for node:=list.Back();node!=nil;node=node.Prev() {
			str=str+node.Value.(string)
			//fmt.Printf("%s",node.Value.(string))
			//fmt.Println(result)
		}
		*result=append(*result,str)

		//fmt.Println("")
		return
	}
	temp:=string(int(chars[i]-'0'-1)+'A')
	list.PushBack(temp)
	translate2(chars,list,i-1,result)
	node:=list.Back()
	list.Remove(node)
	if i-1>=0{
		temp2:=int(chars[i-1]-'0')*10+int(chars[i]-'0')
		if temp2<=26 {
			temp=string(temp2-1+'A')
			list.PushBack(temp)
			translate2(chars,list,i-2,result)
			node=list.Back()
			list.Remove(node)
		}

	}

}
