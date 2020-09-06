package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s,p string
	scan:=bufio.NewScanner(os.Stdin)
	for scan.Scan(){
		s=scan.Text()
		break
	}
	for scan.Scan(){
		p=scan.Text()
		break
	}
	next:=getnext2(p)
	fmt.Println(next)
	ss:=[]byte(s)
	ps:=[]byte(p)
	var i,j int
	for i<len(ss){
		if j==len(p){
			fmt.Println(i-j,i-1,string(ss[i-j]),string(ss[i-1]))
			break
		}
		if j==-1||ss[i]==ps[j]{
			i++
			j++
		}else {
				j=next[j]
		}
	}
}
func getnext(pattern string) []int {
	patterns:=[]byte(pattern)
	var next=make([]int,len(patterns))
	next[0]=-1
	var i,j=0,-1
	for i<len(patterns)-1{
		if j==-1|| patterns[i]==patterns[j]{
			i++
			j++
			next[i]=j
			//fmt.Println(next,i)
		}else {
			j=next[j]
		}
	}

	return next
}
func getnext2(pattern string)[]int  {
	patterns:=[]byte(pattern)
	var next=make([]int,len(patterns))
	next[0]=-1
	var i,j=1,0
	for i<len(patterns)-1{
		if j==-1 || patterns[i]==patterns[j]{
			i++
			j++
			if patterns[i]!=patterns[j]{
				next[i]=j
			}else {
				next[i]=next[j]
			}
		}else {
			j=next[j]
		}
	}
	return next
}
