package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//方法1,直接用+=操作符
	//golang里的字符串是不可变的，每次运算都会产生一个新的字符串，所以产生很多临时的无用的字符串
	//不仅没有用，还会给gc带来额外的负担，所以性能比较差
	var s="hello"
	var str string
	for i:=0;i<10;i++{
		str=s+"--"+str
	}
	fmt.Println("function1",str)

	//方法2 根据formate参数生成格式化的字符串并返回该字符串
	//内部使用[]byte实现，不像直接运算符这种会产生很多临时的字符串，但是内部的逻辑
	//比较复杂，有很多额外的判断，还用到了interface，所以性能也不是很好
	str=fmt.Sprintf("%s--%s",s,s)
	fmt.Println("function2",str)

	//方法3 将一系列字符串连接成一个字符串，之间用sep来分隔
	//join会先根据字符串数组的内容，计算出一个拼接之后的长度，然后申请对应大小的内存
	//一个一个字符串填入，在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小
	var strs []string
	for i:=0;i<10;i++{
		strs=append(strs,s)
	}
	str=strings.Join(strs,"--")
	fmt.Println("function3",str)

	//方法4 利用buffer(buffer是一个实现了读写方法的可变大小的字节缓冲)，将所有的字符串都写入到一个buffer缓冲区中，然后再统一输出，这种方法最快
	//这个比较理想，可以当成可变字符使用，对内存的增长也有优化，如果能预估字符串的长度，还可以使用buffer.Grow接口来设置capacity
	var buffer bytes.Buffer
	for i:=0;i<10;i++{
		buffer.WriteString(s)
	}
	str=buffer.String()
	fmt.Println("function4",str)

}
