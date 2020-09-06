package main

import (
	"fmt"
	"math"
)

func main() {
	digitAtIndex(13,1)
	digitAtIndex2(13)
}
func digitAtIndex(n int,digit int)  {
	if n<0{
		return
	}
	count :=countOfIntegers(digit)
	if n>count { //如果n大于该位数所占的总位数，则往更高的位数去查找
		digitAtIndex(n-count,digit+1)
	}else {//确定n所在的位数
		num1 := n/digit
		num2 := n%digit
		num := num1+int(math.Pow10(digit-1))//确定n所在的数字
		if num2==0{
			fmt.Println((num-1)%10)
			return
		}
		for i:=0;i<digit-num2-1;i++{ //确定n所在数字中的哪一个
			num=num/10
		}
		fmt.Println(num%10)
	}
}
//统计该位数包含多少个数字，并返回这些数字所占的总位数
func countOfIntegers(digit int) int {
	if digit==1 {
		return 10
	}
	count :=int(math.Pow10(digit)-math.Pow10(digit-1))*digit
	return count
}

func digitAtIndex2(n int){
	var count=1
	n=n-10
	for n>0{
		n=n-9*int(math.Pow(10,float64(count)))*(count+1)
		count++
	}
	if count>1{
		n=n+9*int(math.Pow(10,float64(count-1)))*(count)
	}else {
		n=n+10
		fmt.Println(n-1)
		return
	}
	num1:=n/count
	num2:=n%count
	num:=int(math.Pow(10,float64(count-1)))
	num=num+num1-1
	if num2==0{
		fmt.Println(num%10)
	}else {
		num=num+1
		for i:=0;i<count-num2;i++{
			num=num/10
		}
		fmt.Println(num%10)
	}


}

