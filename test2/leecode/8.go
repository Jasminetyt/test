package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("9223372036854775808"))
}
func myAtoi(str string) int {
	if str==""{
		return 0
	}
	chars := []byte(str)
	if chars[0]!=' ' && chars[0]<'0' && chars[0]>'9' && chars[0]!='+' && chars[0]!='-'{
		return 0
	}
	var i int
	for i=0;i<len(chars);{
		if chars[i]==' '{
			i++
		}else {
				break
		}
	}
	if i>=len(chars){
		return 0
	}
	var flag byte
	if chars[i]=='+' || chars[i]=='-' {
		flag=chars[i]
		i++
	}
	var num int
	for i<len(chars)&&chars[i]>='0'&&chars[i]<='9'{
		num=int(chars[i]-'0')+num*10
		i++
		if num> 1<<31{
			break
		}
	}
	switch flag {
	case '-':
		num=-num;
	case '+':
		num=+num;
	default:
		num=num
	}
	if num>math.MaxInt32 {
		return math.MaxInt32
	}else if num<math.MinInt32 {
		return math.MinInt32
	}else {
		return num
	}
}