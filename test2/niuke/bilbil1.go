package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	var nums []int
	var count,sum int
	for scan.Scan(){
		if scan.Text()=="," {
			count=1
		}else if count==1 {
			sum,_ = strconv.Atoi(scan.Text())
			break
		}else {
			num,_ := strconv.Atoi(scan.Text())
			nums=append(nums,num)
		}
	}
	var flag=make([]int,len(nums))
	if isEqualN(sum,3,nums,flag) {
		fmt.Println("True")
	}else {
		fmt.Println("False")
	}

}
func isEqualN(sum ,n int,num ,flag []int) bool {
	if n==0 && sum==0 {
		return true
	}
	if n>0 {
		n--
		for i:=0;i<len(num);i++{
			if flag[i]==0 {
				sum2:=sum-num[i]
				flag[i]=1
				if 	isEqualN(sum2,n,num,flag){
					return true
				}
				flag[i]=0
			}
		}
	}
	return false
}
