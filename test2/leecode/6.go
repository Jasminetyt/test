package main

import "fmt"

func main() {
	var s="LEETCODEISHIRING"
	var numRows=4
	fmt.Println(convert(s,numRows))
}
func convert(s string, numRows int) string {
	if numRows==1{
		return s
	}
	var long []byte
	var short []byte
	var count1 =numRows
	var count2 int
	for i:=0;i<len(s);i++{
		if count1>0 {
			long=append(long,s[i])
			count1--
			if count1==0 {
				count2=numRows-2
				continue
			}
		}

		//fmt.Println("long",string(long),count1)
		if count2>0 {
			short=append(short,s[i])
			count2--
			if count2==0 {
				count1=numRows
			}
		}
		//fmt.Println("short",string(short))
		if count1==0&&count2==0{
			count1=numRows
			long=append(long,s[i])
		}
	}
	//fmt.Println(string(long),string(short))
	var flag=false
	var str string
	for i:=0;i<numRows;i++{
		for j,k:=i,numRows-2-i;j<len(long);j=j+numRows{
			str=str+string(long[j])
			if k>=0&&k<len(short)&&flag{
				str=str+string(short[k])
				k=k+numRows-2
			}
		}
		//fmt.Println(str)
		flag=true
	}
	return str
}