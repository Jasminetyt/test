package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(7,-3))
}
func divide(dividend int, divisor int) int {
	var temp int
	var res int64
	var c= dividend^divisor //使用位运算判断两数的符号是否相同
	dividend=int(math.Abs(float64(dividend)))//取绝对值
	divisor=int(math.Abs(float64(divisor)))
	for dividend >= divisor{
		temp=divisor        //若除数大于加倍后的除数，则加倍后的除数还原至初始除数
		i := 1              //倍数也进行初始化
		for dividend >= temp{
			dividend=dividend-temp//被除数减去除数，
			res=res+int64(i) //记录倍数总和
			i <<= 1 //记录倍数
			temp <<= 1 //除数加倍
		}
	}
	if c<0 {
		res=res*-1
	}
	if res>math.MaxInt32{
		return math.MaxInt32
	}
	if res<math.MinInt32{
		return math.MinInt32
	}
	return int(res)
}