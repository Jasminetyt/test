package main

import (
	"fmt"
	"math"
)

func main() {
	var s string
	var k int
	fmt.Scan(&k,&s)
	chars:=[]byte(s)
	var count,max int
	var mchars=make(map[byte]int)
	var nums []byte
	for i:=0;i<len(chars);i++{
		if _,ok:=mchars[chars[i]];!ok{
			if k==0{
				//fmt.Println("k==0")
				temp:=delete1(mchars)
				count=i-mchars[temp]
				delete(mchars,temp)
				mchars[chars[i]]=i

			}else {
				k--
				mchars[chars[i]]=i
				nums=append(nums,chars[i])
				count++
			}
		}else {
			count++
			mchars[chars[i]]=i
		}
		if count>max{
			max=count
		}
		//fmt.Println("max",max,"count",count)
	}
	fmt.Println(max)
}
func delete1(mchars map[byte]int) byte {
	var min=math.MaxInt32
	var index byte
	for key,value:=range mchars{
		if value<min{
			min=value
			index=key
		}
	}
	return index
}
