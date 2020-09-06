package main

import (
	"fmt"
	"strings"
)

func main() {
	s:="bba"
	t:="ab"
	fmt.Println(minWindow(s,t))
}
func minWindow(s string, t string) string {
	if s==""||t=="" || len(s)<len(t){
		return ""
	}
	var smap=make(map[byte]int)
	for i:=0;i<len(t);i++{
		if n,ok:=smap[t[i]];ok{
			smap[t[i]]=n+1
		}else {
			smap[t[i]]=1
		}
	}
	var min,max,count= 0,0,0
	for !IsContains(s[min],smap){
		min++
	}
	fmt.Println(min,max)
	for ;count<len(t)&&max<len(s);{
		n:=smap[t[count]]
		if strings.Contains(s[min:max+1],string(t[count])){
			smap[t[count]]=0
			count=count+n
			fmt.Println(count)
			for  n>1{
				for max<len(s)&&s[max]!=t[count]{
					max++
				}
				n=n-1
			}
		}else {
			max++
		}
		fmt.Println(count,max)
	}
	fmt.Println(min,max)
	fmt.Println(s[min:max+1])
	if count<len(t) {
		return ""
	}
	var result string
	left,right := min,max+1
	for right<len(s){
		if n,ok :=smap[s[right]];ok{
			smap[s[right]]=n+1
		}
		for IsContains(s[left],smap){
			left++
		}
		if right-left<max-min {
			max=right
			min=left
		}
		right++
	}
	result=s[min:max+1]
	return result
}
func IsContains(char byte,smap map[byte]int) bool {
	if n,ok:=smap[char];!ok{
		return true
	}else {
		if n>0 {
			smap[char]=n-1
			return true
		}
	}
	return false
}


