package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi","issip"))
}
func strStr(haystack string, needle string) int {
	if haystack!=""&&needle=="" {
		return 0
	}
	var haystackbyte=[]byte(haystack)
	var needlebyte=[]byte(needle)
	var flag,j int
	for i:=0;i<len(haystackbyte);{
		if haystackbyte[i]==needlebyte[j]{
			j++
			i++
		}else{
			j=0
			i=flag+1
			flag++
		}
		if j==len(needlebyte){
			break
		}
	}
	if j==len(needlebyte) {
		return flag
	}
	return -1
}