package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	var scan=bufio.NewScanner(os.Stdin)
	for scan.Scan(){
		str=scan.Text()
		break
	}
	var strs []string
	var s string
	for i:=0;i<len(str);{
	if str[i]!=' '{
		s=s+string(str[i])
		i++
	}else {
			if s!=""{
				strs=append(strs,s)
			}
			s=""
			i++
	}
	}
	strs=append(strs,s)
	var result string
	for i:=len(strs)-1;i>=0;i--{
		result=result+" "+strs[i]
	}
	fmt.Println(string(result[1:]))
}
