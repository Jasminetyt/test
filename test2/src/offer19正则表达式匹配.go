package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str ,pattern string
	scan:=bufio.NewScanner(os.Stdin)
	for scan.Scan(){
		str=scan.Text()
		break
	}
	for scan.Scan(){
		pattern=scan.Text()
		break
	}
	//result := match(str,pattern)
	//fmt.Println(result)
	//fmt.Println(match("",""))
	fmt.Println(match3(str,pattern,0,0))
}

func match(str , patterns string) bool {
chars := []byte(str)
patternch := []byte(patterns)
var i ,j =0,0
	for i<len(chars)&&j<len(patternch){
		if chars[i]==patternch[j] {
			fmt.Println("equal")
			i++
			j++
		}else if patternch[j]=='.' { //跳过当前的‘.’和字符中的当前值
			fmt.Println("pattern is .")
			i++
			j++
		}else if patternch[j]=='*' { //若果匹配模式中的字母为*，则出现在*前面的字符可以出现任意次
			temp := patternch[j-1] //*前面的字符
			j++ //pattern向后移一位
			for i<len(chars)&&chars[i]==temp { //如果chars中同temp相同则向后移一位
				if j<len(patternch)&&patternch[j]==temp { //如果patterns中同temp相同则向后移一位
					j++
				}
				i++
			}
		}else if patternch[j+1]=='*' {
			j=j+2
		}
	}
	if i<len(chars) || j<len(patternch) {
		return false
	}
return true
}

func match3(str,pattern string,x ,y int) bool{
	if str==""&&pattern==""{
		return true
	}
	if str!=""&&pattern==""{
		return false
	}
	var i,j=x,y
	for ;j<len(pattern);{
		if j+1<len(pattern)&&pattern[j+1]!='*'{
			if i<len(str)&&(str[i]==pattern[j]||pattern[j]=='.'){
				i++
				j++
			}else {
					return false
			}
		}else if j+1<len(pattern){
			var flag bool
			if i<len(str)&&(str[i]==pattern[j]||pattern[j]=='.'){
				flag=match3(str,pattern,i+1,j+2)||match3(str,pattern,i+1,j)
			}
			return flag || match3(str,pattern,i,j+2)
		}else {
			if i<len(str)&&(str[i]==pattern[j]||pattern[j]=='.'){
				i++
				j++
			}else {
				return false
			}

		}
	}
	if i==len(str)&&j==len(pattern){
		return true
	}else {
		return false
	}
}