package main

import (
	"fmt"
)

//map是无序的
func main() {
	var str string
	fmt.Scanln(&str)
	var chars=[]byte(str)
	if len(chars)<=0 || chars==nil{
		return
	}
	firstNoRepeatingChar(chars)
	//firstAppearingOnce()
	firstNoRepeatingChar2(chars)

}
//字符串中第一个只出现一次的字符
func firstNoRepeatingChar(chars []byte)  {
	mp := make(map[byte]int)
	var mapindex []byte
	var temp byte
	var index=0
	for i:=0;i<len(chars);i++{
		temp=chars[i]
		if value,exist := mp[temp];!exist{
			mp[temp]=1
			mapindex=append(mapindex,temp)
			index++
		}else {
			mp[temp]=value+1
		}
	}
	for _,value := range mapindex{
		if mp[byte(value)]==1{
			fmt.Printf("%c\n",value)
			return
		}
	}
	fmt.Println("不存在只出现一次的字符")
	}
//字符流中第一个只出现一次的字符
func firstAppearingOnce()  {
	var str string
	var chars []byte
	for   {
		fmt.Scanln(&str)
		chars=[]byte(str)
		if len(chars)<=0 || chars==nil{
			continue
		}
		firstNoRepeatingCharInStream(chars)
	}
}
var (
	m=make(map[byte] int)
	s []byte
	position = 0
	)
func firstNoRepeatingCharInStream(chars []byte)  {
	for _,value := range chars{
		if v,exist := m[byte(value)];!exist{
			m[byte(value)]=1
			s=append(s,byte(value))
		}else {
			m[byte(value)]=1+v
		}
	}
	for position<len(s){
		ch :=byte(s[position])
		if m[ch]==1 {
			break
		}
		position++
	}
	if position<len(s) {
		fmt.Printf("当前字符流中第一个只出现一次的字符是：%c \n",s[position])
	}else {
		fmt.Println("当前字符流中不存在只出现一次的字符")
	}
	fmt.Println("---------------")
}

func firstNoRepeatingChar2(chars []byte){
	var sch=make([]byte,len(chars))
	var mch=make(map[byte]int)
	var count int
	for i:=0;i<len(chars);i++{
		if v,ok:=mch[chars[i]];!ok{
			sch[count]=chars[i]
			count++
			mch[chars[i]]=1
		}else {
			mch[chars[i]]=v+1
		}
	}
	for i:=0;i<len(sch);i++{
		if mch[sch[i]]==1{
			fmt.Println(string(sch[i]))
			return
		}
	}
	fmt.Println("not exist")
}





