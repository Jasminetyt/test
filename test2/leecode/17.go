package main

import "fmt"

func main() {
	var digits=""
	fmt.Println(letterCombinations(digits))
}
func letterCombinations(digits string) []string {
	var result []string
	if digits=="" {
		return result
	}
	temp:=make([]byte,len(digits))
	nums:=[]byte(digits)
	digitToapla(&result,0,nums,temp)
	return result
}
func digitToapla(result *[]string,count int,nums,temp []byte)  {
	var chars=map[byte]string{'2':"abc",'3':"def",'4':"ghi",'5':"jkl",'6':"mno",'7':"pqrs",'8':"tuv",'9':"wxyz"}
	if count==len(temp){
		*result=append(*result,string(temp))
	}else {
		if t,exist:=chars[nums[count]];exist{
			ch:=[]byte(t)
			for i:=0;i<len(ch);i++{
				temp[count]=ch[i]
				digitToapla(result,count+1,nums,temp)
			}
		}

	}
}