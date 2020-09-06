package main

import "fmt"


func main() {
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt2("MCMXCIV"))
}
func romanToInt(s string) int {
	var flag=false
	var chars=[]byte(s)
	var num,temp int
	for i:=0;i<len(chars);{
		if chars[i]=='M' {
			temp,flag=getnum('M',flag)
			num=num+temp
			i++
		}else if chars[i]=='D' {
			temp,flag=getnum('D',flag)
			num=num+temp
			i++
		}else if chars[i]=='C' {
			if i+1<len(chars)&&(chars[i+1]=='D'||chars[i+1]=='M') {
				flag=true
			}else {
				temp,flag=getnum('C',flag)
				num=num+temp

			}
			i++
		}else if chars[i]=='L' {
			temp,flag=getnum('L',flag)
			num=num+temp
			i++
		}else if chars[i]=='X' {
			if i+1<len(chars)&&(chars[i+1]=='L'||chars[i+1]=='C') {
				flag=true
			}else {
				temp,flag=getnum('X',flag)
				num=num+temp
			}
			i++
		}else if chars[i]=='V' {
			temp,flag=getnum('V',flag)
			num=num+temp
			i++
		}else if chars[i]=='I' {
			if i+1<len(chars)&&(chars[i+1]=='V'||chars[i+1]=='X' ){
				flag=true
			}else {
				temp,flag=getnum('I',flag)
				num=num+temp
			}
			i++
		}
	}
	return num
}
var nums1=map[byte]int{'M':1000,'D':500,'C':100,'L':50,'X':10,'V':5,'I':1}
var nums2=map[byte]int{'M':900,'D':400,'C':90,'L':40,'X':9,'V':4,}
func getnum(char byte,flag bool)(int,bool){
	if flag {
		flag=false
		return nums2[char],flag
	}else {
		return nums1[char],flag
	}
}
func romanToInt2(s string) int {
	var nums=map[string]int{"I":1,"IV":4,"V":5,"IX":9,"X":10,"XL":40,"L":50,"XC":90,"C":100,"CD":400,"D":500,"CM":900,"M":1000}
	var chars=[]byte(s)
	var result int
	for i:=0;i<len(chars);{
		if i+1<len(chars) {
			str:=fmt.Sprintf("%s%s",string(chars[i]),string(chars[i+1]))
			if num, exist:=nums[str];exist{
				result=result+num
				i=i+2
			}else {
				result=result+nums[string(chars[i])]
				i++
			}
		}else {
			result=result+nums[string(chars[i])]
			i++
		}
	}
	return result
}