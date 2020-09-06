package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	var n int
	var s string
	var result []string
	for i:=0;i<t;i++{
		fmt.Scan(&n)
		fmt.Scan(&s)
		if n<11{
			result=append(result,"NO")
		}else {
			ss:=[]byte(s)
			var i int
			for i=0;i<len(ss);i++{
				if ss[i]=='8'{
					break
				}
			}
			if len(ss)-i>=11{
				result=append(result,"YES")
			}else {
				result=append(result,"NO")
			}
		}
	}
	for i:=0;i<len(result);i++{
		fmt.Println(result[i])
	}
}
