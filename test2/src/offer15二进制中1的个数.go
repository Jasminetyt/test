package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var count=0
	//方法1
	var flag=1
	for flag > 0{
		if n & flag > 0{
			count++
		}
		flag=flag>>1
	}
	fmt.Println(count)

	//方法2
	//count=0
	//for n > 0 {
	//	count++
	//	n=(n-1) & n;
	//}
	//fmt.Println(count)
}
