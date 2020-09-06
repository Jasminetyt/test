package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
//维护一个数组MaxV[i]，记录长度为i的递增子序列中最大元素的最小值，
// 并对于数组中的每个元素考察其是哪个子序列的最大元素，二分更新MaxV数组，最终i的值便是最长递增子序列的长度。
//时间复杂度为O(NlogN) 空间复杂度为O(N)
func main() {
	scan:=bufio.NewScanner(os.Stdin)
	var str string
	for scan.Scan(){
		str=scan.Text()
		break
	}
	strs:=strings.Split(str," ")
	var num=make([]int,len(strs))
	for i:=0;i<len(strs);i++{
		num[i],_=strconv.Atoi(strs[i])
	}

	fmt.Println(lis(num))
}
 func lis(num []int) int {
 	var maxmin []int
 	maxmin=append(maxmin,num[0])
 	var temp=0
 	for i:=1;i<len(num);i++{
 		if num[i]>maxmin[temp]{
 			maxmin=append(maxmin,num[i])
 			temp++
		}else{
 				index:=findmaxmin(maxmin,num[i])
 				maxmin[index]=num[i]
		}
	}
 	return temp+1
 }

func findmaxmin(maxmin []int,x int) int {
	left,right:=0,len(maxmin)-1
	for left<=right{
		mid := (left+right)/2
		if maxmin[mid]<=x{
			left=mid+1
		}else {
			right=mid-1
		}
	}
	return left
}
