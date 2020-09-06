package main

import "fmt"

func main() {
	fmt.Println(getUglyNumber(1500))
	getUglyNumber2(1500)
}
func getUglyNumber(index int) int {
	if index <= 0 {
		return 0
	}
	var nums=make([]int,index)
	var min int
	nums[0]=1
	m2 , m3 , m5,i :=0,0,0,0
	for i=1;i<index;i++{
		max :=nums[i-1]
		for m2<i&&nums[m2]*2<=max{
			m2++
		}
		for m3<i&&nums[m3]*3<=max{
			m3++
		}
		if nums[m2]*2>nums[m3]*3 {
			min=nums[m3]*3
		}else {
			min=nums[m2]*2
		}
		for m5<i&&nums[m5]*5<=max{
			m5++
		}
		if nums[m5]*5<min {
			min=nums[m5]*5
		}
		nums[i]=min
	}
	return nums[i-1]
}
func getUglyNumber2(index int){
	var index1,index2,index3=0,0,0
	var num=make([]int,index)
	num[0]=1
	for i:=1;i<index; {
		temp1:=num[index1]*2
		temp2:=num[index2]*3
		temp3:=num[index3]*5
		if temp1<=temp2&&temp1<=temp3{
			if temp1>num[i-1]{
				num[i]=temp1
				i++
			}
			index1++
		}else if temp2<=temp3&&temp2<temp1{
			if temp2>num[i-1]{
				num[i]=temp2
				i++
			}
				index2++
		}else {
			if temp3>num[i-1]{
				num[i]=temp3
				i++
			}
					index3++
		}
	}
	fmt.Println(num[len(num)-1])

}