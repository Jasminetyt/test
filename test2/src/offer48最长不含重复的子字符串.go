package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	var chars =[]byte(str)
	if len(chars)==0 || chars==nil {
		return
	}
	fmt.Println(longestSubstringWithoutDuplication(chars))
	longestSubstringWithoutDuplication3(chars)
}
func longestSubstringWithoutDuplication(chars []byte) int {
	var mp=make(map[byte]int)
	var maxlength=make([]int,len(chars))
	maxlength[0]=1
	mp[chars[0]]=0
	max := 1
	for i:=1;i<len(chars);i++{
		str := chars[i]
		if value,exist := mp[str];exist{
			if (i-value)>maxlength[i-1] {
				maxlength[i]=maxlength[i-1]+1
			}else {
				maxlength[i]=i-value
			}
			mp[str]=i
		}else {
			mp[str]=i
			maxlength[i]=maxlength[i-1]+1
		}
		if max<maxlength[i] {
			max=maxlength[i]
		}
	}
	return max
}

func longestSubstringWithoutDuplication3(chars []byte){
	var mmap=make(map [byte]int)
	var start,max=0,0
	for i:=0;i<len(chars);i++{
		if v,ok:=mmap[chars[i]];!ok{
			mmap[chars[i]]=i
		}else {
			mmap[chars[i]]=i
			if v>=start{
				start=v+1
			}
		}
		if i-start+1>max{
			max=i-start+1
		}
	}
	fmt.Println(max)
}
