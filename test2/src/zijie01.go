package main

import "fmt"

func main()  {
	var n int
	fmt.Scanln(&n)
	var numberk=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&numberk[i])
	}
	var mapk=make(map[int][]int)
	for i:=0;i<n;i++{
		key := numberk[i]
			mapk[key]=append(mapk[key],i+1)
	}
	var count,l,r,k ,number int
	fmt.Scan(&count)
	var result=make([]int,count)
	for i:=0;i<count;i++{
		fmt.Scan(&l,&r,&k)
		temp:=mapk[k]
		for j:=0;j<len(temp);j++{
			if temp[j]>=l&&temp[j]<=r {
				number++
			}
		}
		result[i]=number
		number=0
	}
	for i:=0;i<len(result);i++{
		fmt.Println(result[i])
	}
}