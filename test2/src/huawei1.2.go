package main

import "fmt"

func main() {
	var str string
	fmt.Scanln(&str)
	var chars =[]byte(str)
	if len(chars)==0 || chars==nil {
		return
	}
	max,temp:=longestSubstringWithoutDuplication2(chars)
	var str2 string
	for i:=0;i<max;i++{
		str2=string(chars[temp-i])+str2
	}
	fmt.Println(str2)
}
func longestSubstringWithoutDuplication2(chars []byte) (int,int) {
	var mp=make(map[byte]int)
	var maxlength=make([]int,len(chars))
	maxlength[0]=1
	mp[chars[0]]=0
	max := 1
	var temp =0
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
			temp=i
		}
	}
	return max,temp
}

