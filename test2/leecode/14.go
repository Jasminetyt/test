package main

import "fmt"

func main() {
	var strs=[]string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	if len(strs)==0 || strs[0]=="" {
		return ""
	}
	var s string
	n:=0 //控制单个字符串中字符的位置
	for   {
		if n>=len(strs[0]) {
			break
		}
		i:=1 //控制字符串数组中字符串的位置
		for ;i<len(strs);i++{//先将字符串数组中每个字符串的第一个字符进行比较
			if len(strs[i])<=n || strs[0][n]!=strs[i][n] {
				break
			}
		}
		if len(strs)>i {
			break
		}else {
			s=strs[0][:n+1]
			n++
		}
	}
	return s
}