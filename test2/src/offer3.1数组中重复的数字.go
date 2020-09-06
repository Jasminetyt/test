package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
//输入操作开始
	scan := bufio.NewScanner(os.Stdin)
	var count = 0
	var str string
	for scan.Scan() { //输入
		str = scan.Text()
		count++
		if count==1 {
			break;
		}
	}
    strs:=strings.Split(str," ") //去空格
    var num []int
    var temp int
    for i:=0;i<len(strs);i++{ //转为整型
    	temp,_=strconv.Atoi(strs[i])
    	num=append(num,temp)
	}
    //输入操作结束

	var result []int
	for i:=0;i<len(strs);{
		if i==num[i]{
			i++
		}else {
				temp:=num[i]
				if num[temp]==temp{
					result=append(result,temp)
					i++
				}else {
					num[i],num[temp]=num[temp],num[i]
				}
		}
	}
	fmt.Println(result)
}