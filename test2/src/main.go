package main

import (
	"fmt"
)

func main()  {
arr := []int{3,4,3,1,4,5,6,2}
var result []int
result=mergeSort(arr)
	fmt.Printf("result=%v \n",result)
	chars := []byte("sdfgdgd")
	//for i ,v := range chars{
	//	fmt.Printf("i=%d ,v=%c \n",i,v)
	//}
	fmt.Printf("chars=%c\n",chars)
	fmt.Printf("chars=%v\n",chars)
	s := string(chars)
	fmt.Println(s)
}

func mergeSort(arr []int) []int {
	var length=len(arr)
	middle := length/2
	if length >= 2 {
		left := arr[0:middle]
		right := arr[middle :]
		return merge(mergeSort(left) ,mergeSort(right))
	}else {
		return arr;
	}
	
	
}

func merge(left []int , right []int) []int {
	length1 := len(left)
	length2 := len(right)
	result := make([]int,length2+length1)
	var index=0
	i,j := 0,0
	for ; i<length1&&j<length2;{
		if left[i]>right[j] {
			result[index]=right[j]
			j++;
			index++;
		}else {
			result[index]=left[i]
			i++
			index++
		}
	}
	for i<length1 {
		result[index]=left[i];
		i++;
		index++;
	}
	for j<length2 {
		result[index]=right[j];
		j++;
		index++;
	}
	return result
}
