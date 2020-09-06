package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	chars :=[]byte(s)
	var result []string
	var ls =list.New()
	var temp string

	for i:=0;i<len(chars);i++{
		if chars[i]=='['{
			i++
			ls.PushFront(string(chars[i]))
			//fmt.Println(ls.Front().Value.(string))
		}else if ls.Len()==0{
			result=append(result,string(chars[i]))
			//fmt.Println(result)
		}else if chars[i]!=']'{
			ls.PushFront(string(chars[i]))
			//fmt.Println(ls.Front().Value.(string))
		}else {
			var list *list.Element
			list=ls.Front()
			for ;list.Value.(string)!="|";{
				temp=string(list.Value.(string))+temp
				ls.Remove(list)
				list=ls.Front()
			}
			ls.Remove(list)
			temp3:=ls.Front()
			n,_:=strconv.Atoi(temp3.Value.(string))
			//fmt.Println(temp,n)
			var temp2 string
			for i:=0;i<int(n);i++{
				temp2=temp2+temp
			}
			temp=""
		//	fmt.Println("temp2",temp2)
			ls.Remove(temp3)
			//fmt.Println(ls.Len())
			if ls.Len()!=0 {
				temp4:=ls.Front()
				if temp4.Value!="|"{
					temp2=temp4.Value.(string)+temp2
					fmt.Println(temp2)
					ls.Remove(temp4)
					ls.PushFront(temp2)
				}else {
					ls.PushFront(temp2)
				}
			//	fmt.Println(ls.Front().Value.(string))
			}else {
				//fmt.Println(temp2)
				result=append(result,temp2)
			//	fmt.Println(result)
			}
		}
	}
	var rs string
	for i:=0;i<len(result);i++{
		rs=rs+result[i]
	}
	fmt.Println(rs)
}
