package main

import (
	"fmt"
	"strconv"
)
//负-负
//正-正
//正-负
//负-正
func main() {//现在情况为正的大数-正的小数 即正-正
	var str1 string
	var str2 string
	fmt.Scan(&str1)
	fmt.Scan(&str2)
	bigDigitSub(str1,str2)
}
func bigDigitSub(str1,str2 string)  {
	char1 := []byte(str1)
	char2 := []byte(str2)
	char1,char2=comapre(char1,char2) //将较大的数放前面
	result := sub(char1,char2)
	fmt.Println(result)


}
func comapre(char1,char2 []byte) ([]byte,[]byte) {
	length1 ,length2 := len(char1),len(char2)
	if length1 >length2 {
		return char1,char2
	}else if length1<length2 {
		return char2,char1
	}else {
		var i int
		for i=0;i<length2;i++ {
			if char1[i]!=char2[i] {
				break
			}
		}
		if char2[i]>char1[i] {
			return char2[i:],char1[i:]
		}else {
			return char1[i:],char2[i:]
		}
	}
}

func sub(char1,char2 []byte) string {
var (
	i=len(char1)-1
	j=len(char2)-1
	flag=0
	temp int
	str=""
	)
	for j>=0  {
		temp=int(char1[i])-int(char2[j])-flag
		if temp<0 {
			temp=temp+10
			flag=1
		}else {
			flag=0
		}

		if i==0 && temp==0{ //当最后一位为0时，舍弃（char1和char2长度相同时）
			i-- //对最后一位操作后，需要及时减一，否则下面的for循环还要继续操作
			j--
			break
		}
		str=strconv.Itoa(temp)+str
		temp=0
		i--
		j--
	}
	for i>=0 {
		if flag >0{
			temp=int(char1[i]-'0')-flag
			if i==0 && temp==0{
			break
			}
			str=strconv.Itoa(temp)+str
			flag=0
		}else {
			str=string(char1[i])+str
			}
		i--
	}
return str
}