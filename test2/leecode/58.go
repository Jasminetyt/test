package main

import "fmt"

func main() {
	var s=" h  llo  n  "
	fmt.Println(lengthOfLastWord(s))
}
func lengthOfLastWord(s string) int {
	var count int
	var flag=false
	for i:=len(s)-1;i>=0;i--{
		if s[i]==' '{
			if flag{
				break
			}
		}else {
			count++
			flag=true
		}

	}
	return count
}