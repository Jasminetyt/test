package main

import "fmt"

func main() {
	var str string
	var n int
	fmt.Scan(&str,&n)
	var chars=[]byte(str)
	var i int
	for i=0;i<len(chars);i=i+n{
		if i+n<len(chars){
			fmt.Println(string(chars[i:i+n]))
		}else {
			break
		}
	}
	var s =string(chars[i:len(chars)])
	for i=len(s);i<n ; i++ {
		s=s+"0"
	}
	fmt.Println(s)

}
