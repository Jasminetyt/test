package main

import "fmt"

func main() {
	fmt.Println(isMatch("aaaaaaaaaaaaab","a*a*a*a*a*a*a*a*a*a*c"))
}
func isMatch(s string, p string) bool {
	if s==""&& p=="" {
		return true
	}
	if s!="" && p==""{
		return false
	}
	ss:=[]byte(s)
	ps:=[]byte(p)
	var s1,p2,p1 string
	s1=isexit(ss,1)
	p1=isexit(ps,1)
	p2=isexit(ps,2)
		if len(ps)>1&&ps[1]=='*' {//当第二位是*时
			if len(ss)>0&&(ps[0]==ss[0] || ps[0]=='.' ){//若模式和字符串的第一位相等
				return isMatch(s1,p2)||isMatch(s1,string(ps))||isMatch(string(ss),p2)
				//将第一位匹配的视为1次 || 第一位匹配了很多次 || 将第一位匹配的视为0次
			}else {//若模式和字符串的第一位不相等，则字符串不变，模式后移两位(表示模式与字符串匹配0次)
				return isMatch(string(ss),p2)
			}
		}
		if len(ss)>0&&(ss[0]==ps[0] || ps[0]=='.'){
			return isMatch(s1,p1) //当第二位不是*，只需判断模式和字符串的第一位是否相等
		}
	return false
}
func isexit(num []byte,n int) string {
	if len(num)>n {
		return string(num[n:])
	}else {
		return ""
	}
}