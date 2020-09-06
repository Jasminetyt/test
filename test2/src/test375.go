package main

import "fmt"

func main()  {
	fmt.Println(getMoneyAmount(20))
}
func getMoneyAmount(n int) int {
	if n==1 {
		return 0
	}
	return get(1,n)
}

func get(left ,right int) int  {
	if left == right{
		return 0
	}
	if left+1==right {
		return left
	}
	mid := right-3
	if mid<left {
		return right-1
	}
	var temp int
	if get(left,mid-1)>get(mid+1,right) {
		temp = get(left,mid-1)
	}else {
		temp = get(mid+1,right)
	}
	return mid+temp
}

