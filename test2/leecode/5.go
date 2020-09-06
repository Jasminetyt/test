package main

import (
	"fmt"
	"math"
)

func main() {
	var s="babad"
	fmt.Println(longestPalindrome(s))
}
func longestPalindrome(s string) string {
	if s=="" || len(s)<0{
		return ""
	}

	var start,end int
	for i:=0;i<len(s);i++{
		len1 := expandCenter(s,i,i)
		len2 := expandCenter(s,i,i+1)
		length := int(math.Max(float64(len1),float64(len2)))
		if length>end-start{
			start=i-(length-1)/2
			end=i+length/2
		}
	}
	return s[start:end+1]
}

func expandCenter(s string,left,right int) int {
	L,R := left,right
	for L>=0 && R<len(s)&&s[L]==s[R]{
		L--
		R++
	}
	return R-L-1
}