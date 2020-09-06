package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int
	fmt.Scanln(&sum)
	var nums []int
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	for scan.Scan(){
		if scan.Text()=="end" {
			break
		}
		num,_ := strconv.Atoi(scan.Text())
		nums=append(nums,num)
	}
	findNumbersWithSum(nums,sum)
	findContinuousSequence(sum)
}
func findNumbersWithSum(nums []int,sum int)  {
	if len(nums)<0 || nums==nil{
		return
	}
	for i,j:=0,len(nums)-1;i < j ;  {
		result := nums[i]+nums[j]
		if result==sum {
			fmt.Printf("%d + %d = %d",nums[i],nums[j],sum)
			return
		}
		if result<sum {
			i++
		}
		if result>sum {
			j--
		}
	}
	fmt.Printf("不存在两个数和为%d",sum)
}
func findContinuousSequence(sum int)  {
	if sum<=2 {
		return
	}
	small,big,end := 1,2,(sum+1)/2
	result := small+big
	for small < end{
		if result==sum{
			fmt.Printf("%d--%d 的和为%d \n",small,big,sum)
			result=result-small
			small++
		}else if result<sum {
			big++
			result=result+big
		}else if result>sum {
			result=result-small
			small++
		}
	}
}