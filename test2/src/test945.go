package main

import (
	"fmt"
	"sort"
)

func main()  {
	arr := []int{3,2,1,2,1,7}
result := minIncrementForUnique(arr)
fmt.Printf("The count of move=%d",result)
}
func minIncrementForUnique(A []int) int {
	var n =0
    sort.Ints(A)
	for i := 0;i< len(A)-1 ;i++  {
		if A[i]==A[i+1] {
			A[i+1]++
			n++
		}else if A[i]>A[i+1] {
			n=n+A[i]+1-A[i+1]
			A[i+1]=A[i]+1
		}
	}
	return n
}
