package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	scan:=bufio.NewScanner(os.Stdin)
	for scan.Scan(){
		str=scan.Text()
		break
	}
	strs:=strings.Split(str," ")
	var result []string
	for i:=0;i<len(strs);i++{
		if isright(strs[i])!=""{
			result=append(result,isright(strs[i]))
			//fmt.Println(result)
		}
	}
	//for i:=len(result)-1;i>0;i--{
	//	fmt.Println(result[i])
	//}
	for i:=len(result)-1;i>0;i--{
		fmt.Printf("%s ",result[i])
	}
	fmt.Printf("%s",result[0])
}
func isright(str string) string {
	//fmt.Println(str)
	var result string
	var flag bool
	for i:=0;i<len(str);{
		if (str[i]>='a'&&str[i]<='z')||(str[i]>='A'&&str[i]<='Z')||(str[i]>='0'&&str[i]<='9'){
			result=result+string(str[i])
			flag=true
			i++
		}else if str[i]=='-'&&flag&&i+1<len(str)&&str[i+1]!='-'{
			result=result+string(str[i])
			i++
		}else if str[i]=='-'&&flag&&i+1<len(str){
				//fmt.Println("two -")
				var j=i
				for ;j<len(str)&&str[j]=='-';j++{
				}
				result=result+" "
				i=j
		}else {i++}

	}
	return result
}

