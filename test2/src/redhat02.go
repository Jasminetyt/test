package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)
	var num=make([]int,m)
	for i:=0;i<len(num);i++{
		fmt.Scan(&num[i])
	}
	var result1,count1 int
	result1=num[0]
	count1=1
	for i:=2;i<len(num);i=i+2{
	if i+1<len(num){
		if i+2<len(num){
			if i+3<len(num){
				if num[i]+num[i+2]>num[i+1]+num[i+3]{
					result1=result1+num[i]
					count1=count1+1
				}else {
					result1=result1+num[i+1]
					count1=count1+1
					i++
				}
			}else {
				if num[i]+num[i+2]>num[i+1] {
					result1=result1+num[i]
					count1=count1+1
				}else {
					result1=result1+num[i+1]
					count1=count1+1
					i++
				}

			}
		}else {
			if num[i]>num[i+1]{
				result1=result1+num[i]
				count1=count1+1
			}else {
				result1=result1+num[i+1]
				count1=count1+1
				i++
			}
		}
	}else {
		result1=result1+num[i]
		count1=count1+1
	}
	}
	var result2 ,count2 int
	result2=num[1]
	count2=1
	for i:=3;i<len(num);i=i+2{
		if i+1<len(num){
			if i+2<len(num){
				if i+3<len(num){
					if num[i]+num[i+2]>num[i+1]+num[i+3]{
						result2=result2+num[i]
						count2=count2+1
					}else {
						result2=result2+num[i+1]
						count2=count2+1
						i++
					}
				}else {
					if num[i]+num[i+2]>num[i+1] {
						result2=result2+num[i]
						count2=count2+1
					}else {
						result2=result2+num[i+1]
						count2=count2+1
						i++
					}

				}
			}else {
				if num[i]>num[i+1]{
					result2=result2+num[i]
					count2=count2+1
				}else {
					result2=result2+num[i+1]
					count2=count2+1
					i++
				}
			}
		}else {
			result2=result2+num[i]
			count2=count2+1
		}
	}
	//fmt.Println(result1,count1,result2,count2)
	if result1>result2{
		fmt.Println(result1,count1)
	}else if result2>result1{
		fmt.Println(result2,count2)
	}else {
		if count1<count2 {
			fmt.Println(result1,count1)
		}else {
			fmt.Println(result2,count2)
		}
	}

}
