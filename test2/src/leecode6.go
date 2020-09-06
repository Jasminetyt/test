package main

import "fmt"

func main() {
	convert("LEETCODEISHIRING",3)
}
func convert(s string, numRows int) string {
	if s=="" || numRows<=1{
		return s
	}
	chars := []byte(s)
	length := len(chars)
	var numColums int
	temp := length%(2*numRows-2)
	if temp<=numRows {
		if temp!=0 {
			numColums=1
		}
	}else {
		numColums=temp-numRows+1
	}
	numColums=numColums+length/(2*numRows-2)*(numRows-2+1)
	var char=make([][]byte,numRows)
	for i:=0;i<numRows;i++{
		char[i]=make([]byte,numColums)
	}
	var row,colum =0,-1
	for i:=0;i<length;i++{
		temp=i%(2*numRows-2)
		if temp==0 {
			row=0
			colum++
		}else if temp<numRows{
			row++
		}else {
				colum++
				row--
		}
		copy(char,chars[i],row,colum)
	}
	var str string
	for i:=0;i<numRows;i++{
		for j:=0;j<numColums;j++{
			if char[i][j]!='0' {
				str=str+string(char[i][j])
			}
		}
	}
	fmt.Println(str)
	return str
}
func copy(char [][]byte,b byte,i,j int)  {
	char[i][j]=b
}