package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("aaaaa"))//abba
}
func lengthOfLongestSubstring(s string) int {
	chars := []byte(s)
	maps := make(map[byte]int)
	var left ,right,max =0 ,-1,0
	for i:=0;i<len(chars);i++ {
		value,exist := maps[chars[i]]
		if exist {
			if maps[chars[i]]<left {
				right=i
			}else {
				right++
				left=value+1
			}
			maps[chars[i]]=i
		}else {
			maps[chars[i]]=i
			right++
		}
		if right-left+1>max {
			max=right-left+1
		}
		//fmt.Printf("i=%d,left=%d,riht=%d\n",i,left,right)
	}
		return max
}