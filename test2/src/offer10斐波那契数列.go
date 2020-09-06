package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Fibonacci(n))
}
func Fibonacci(n int) int {
	var num1,num2=0,1
	if n==0{
		return num1
	}
	if n==1{
		return num2
	}
	var temp int
	for i:=2;i<=n;i++{
		temp=num1+num2
		num1=num2
		num2=temp
	}
	return temp
}
