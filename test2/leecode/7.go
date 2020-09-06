package main

import (
	"math"
	"strconv"
)

func main() {
	reverse(-123)
}
func reverse(x int) int {
	var i int
	if x<0 {
		i=1
	}
	str:= strconv.Itoa(x)
	chars:=[]byte(str)
	str=""
	for ;i<len(chars) ;i++  {
		if str==""&&chars[i]=='0'{
			continue
		}
		str=string(chars[i])+str
	}
	if x<0 {
		str="-"+str
	}
	x,_=strconv.Atoi(str)
	if x>math.MaxInt32|| x<math.MinInt32{
		x= 0
	}
	return x
}
func reverse2(x int) int {
	var num int
	for x!=0{
		num=num*10+x%10
		x=x/10
	}
	if num>math.MaxInt32 || num<math.MinInt32{
		num=0
	}
	return num
}