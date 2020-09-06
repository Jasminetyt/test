package main

import (
	"fmt"
)

//当指数小于等于0时
//当底数为0指数小于等于0时,为无效，返回0
//更高效的算法，例如8次方可分为进行3次2次方的运算
//位运算比乘除法及求余运算的效率要高很多
func main() {
	fmt.Println(power(-2,-3))

}
func power(base float64,exponent int) float64 {
	var abs =false
	if base==0.0{
		if exponent<=0 {
			fmt.Println("无效输入")
			return -1.0//无效输入
		}else {
			return 0.0
		}
	}
	if exponent==0 {
		return 1
	}
	if exponent<0 {
		abs = true
		exponent=-exponent
	}
	result := powerWithUnsignedExponent(base,exponent)
	if abs {
		result=1.0/result
	}
	return result
}
func powerWithUnsignedExponent(base float64 , exponent int) float64 {
	if exponent==1 {
		return base
	}
	result := powerWithUnsignedExponent(base,exponent>>1)
	result*=result
	if exponent & 0x1==1{ //&表示按位与计算 0x开头的表示16进制 k & 0x1 其效果为取k的二进制中最右边的数字，该式也可以用作判断k的奇偶性
		result*=base      //如果k为奇数，其计算结果为1，否则为0
	}                     //最后一次返回时才会进行该操作
	return result
}

func power2(base float64,exponent int) float64 {
	if base==0{
		return 0
	}
	var flag=true
	if exponent<0{ //指数小于零，取倒数
		exponent=-exponent
		flag=false
	}
	if flag{
		return powerWithUnsignedExponent2(base,exponent)
	}
	return 1.0/powerWithUnsignedExponent2(base,exponent)
}

func powerWithUnsignedExponent2(base float64 , exponent int) float64{
	if exponent==0{
		return 1
	}
	if exponent==1{
		return base
	}
	result := powerWithUnsignedExponent2(base,exponent>>1)
	result *= result
	if exponent & 0x1==1{ //是否为奇数
		result *= base
	}
	return result
}