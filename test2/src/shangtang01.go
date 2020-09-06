package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var result=1
	var count int
	if n==4{
		result=result*2*2
		count=count+2
	}else if n==3{
		result=result*2*1
		count=count+2
	}else if n==2{
		result=result*1*1
		count=count+2
	}else {
		for n>=5{
			result=result*3
			n=n-3
			count++
		}
		if n==4{
			result=result*2*2
			count=count+2
		}else {
			result=result*n
			count++
		}

	}

	fmt.Println(result+count)

}
