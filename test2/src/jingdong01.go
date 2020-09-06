package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	var str string
	scan :=bufio.NewScanner(os.Stdin)
	for scan.Scan(){
		str=scan.Text()
		break
	}
	chars:=[]byte(str)
	fmt.Println(string(chars))
	var flag=0 //小写
	var count int
	for i:=0;i<n;i++{
		if chars[i]>='a'&&chars[i]<='z'&&flag!=0{
			count++
			flag=0
		}else if chars[i]>='A'&&chars[i]<='Z'&&flag!=1 {
			count++
			flag=1
		}
	}
	fmt.Println(n+2)
}
