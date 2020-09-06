package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(6))
}
func countAndSay(n int) string {
	if n==1{
		return "1"
	}
	var str = make([]string,1)
	str[0]="1"
	for i:=1;i<n;i++{
		count:=1
		s:=""
		chars := []byte(str[i-1])
		temp := chars[0]
		for j:=1;j<len(chars);j++{
			if temp==chars[j]{
				count++
			}else{
				s=s+strconv.Itoa(count)+string(temp)
				temp=chars[j]
				count=1
			}
		}
		s=s+strconv.Itoa(count)+string(temp)
		str=append(str,s)
	}
	return str[len(str)-1]
}
