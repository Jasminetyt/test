package main

import (
	"fmt"
)

func main() {
		var s string
		for {
			a,_:=fmt.Scan(&s)
			if a==0 {
				break
			}else {
				var left,right=5,6
				var temp1,temp2,count int
				chars := []byte(s)
				for i:=0;i<len(chars);i++{
					number :=int(chars[i]-'0')
					if number<=left {
						if left-temp1<=number {
							temp1=temp1-(left-number)
						}else {
							count=left-temp1-number+count
							temp1=0
						}
						left=number
						temp2=temp2+2
						count=count+1
					}else if number>=right {
						if right+temp2>=number {
							temp2=temp2-(number-right)
						}else {
							count=number-(right+temp2)+count
							temp2=0
						}
						right=number
						temp1=temp1+2
						count=count+1
					}else {
						abs1:=number-left
						abs2:=right-number
						if temp2>abs2&&temp1<abs1{
							temp2=temp2-abs2
							temp1=temp1+abs2
						}else if temp1>abs1&&temp2<abs2 {
							temp1=temp1-abs1
							temp2=temp2+abs1
						}else if abs1<=abs2 {
							if left+temp1>=number {
								temp1=temp1-(number-left)
							}else {
								count=number-(left+temp1)+1+count
								temp1=0
							}
							left=number
							temp2=temp2+2
						}else {
							if right-temp2<=number{
								temp2=temp2-(right-number)
							}else {
								count=right-temp2-number+1+count
								temp2=0
							}
							right=number
							temp1=temp1+2
						}
						count=count+1
					}
				}
				fmt.Println(count)
			}
		}

}
