package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"test2/tree"
)

func main() {
scan := bufio.NewScanner(os.Stdin)
scan.Split(bufio.ScanWords)
var nums []int
	for scan.Scan() {
		if scan.Text()=="end" {
			break
		}
		num,_:=strconv.Atoi(scan.Text())
		nums=append(nums,num)
	}
var tr *tree.Tree
result:=tr.VerifySquenceOfBST(nums,0,len(nums)-1)
fmt.Println(result)

fmt.Println(isoutOrder(nums,0,len(nums)-1))
}

func isoutOrder(num []int,start,end int)bool{
	if start>=end{
		return true
	}
	index:=num[end]
	var i,j int
	for i=start;i<end;i++{
		if num[i]>index{
			break
		}
	}
	for j=i;j<end;j++{
		if num[j]<index{
			break
		}
	}
	if j<end{
		return false
	}
	return isoutOrder(num,start,i-1)&&isoutOrder(num,i,end-1)
}
