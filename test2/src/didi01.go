package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	var strs=make([]string,2*n-1)
	for i:=0;i<len(strs);i++{
		fmt.Scan(&strs[i])
	}
	var index1 =0
	for i:=1;i<len(strs); {
		if strs[i]=="+"{
			i=isplus(i+2,strs)
			if strs[i]=="*"||strs[i]=="/"{
				i=i-2
			}else {
				i=i-1
			}
			order(index1,i-1,strs)
		}else if strs[i]=="*"{
			i=ismul(i+2,strs)
			order(index1,i-1,strs)
		}else if strs[i]=="-"||strs[i]=="/"{
			i++
		}
		i=i+2
		index1=i-1
	}
	for i:=0;i<len(strs)-1;i++{
		fmt.Printf("%s ",strs[i])
	}
	fmt.Printf("%s",strs[len(strs)-1])
}
func isplus(i int,strs []string) int {
	if i<len(strs) &&strs[i]=="+"{
	  return isplus(i+2,strs)
	}
	return i
}
func ismul(i int,strs []string) int {
	if i<len(strs)&&strs[i]=="*"{
		return ismul(i+2,strs)
	}
	return i
}
func order(x,y int ,strs []string)  {
	var nums []int
	for i:=x;i<=y;i=i+2{
		temp,_:=strconv.Atoi(strs[i])
		nums=append(nums,temp)
	}
	sort.Ints(nums)
	var count=0
	for i:=x;i<=y;i=i+2{
		strs[i]=strconv.Itoa(nums[count])
		count++
	}
	fmt.Println(strs)
}
