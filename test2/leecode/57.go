package main

import "fmt"

func main() {
	var intervals=[][]int{{1,5}}
	var newInterval=[]int{5,7}
	fmt.Println(insert(intervals,newInterval))
}
func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	no1,no2 := newInterval[0],newInterval[1]
	var i int
	for i=0;i<len(intervals);i++{
		temp :=intervals[i][1]
		if no1<=temp {
			break
		}
		result=append(result,intervals[i])
	}
	var temp2=make([]int,2)
	if i==len(intervals)|| no1<intervals[i][0]{
		temp2[0]=no1
	}else {
		temp2[0]=intervals[i][0]
	}
	var last =no2
	for j:=i;j<len(intervals);{
		if intervals[j][0]<no2&&intervals[j][1]<no2{
			j++
		}else if intervals[j][0]>no2 {
			last=no2
			i=j
			break
		}else if intervals[j][1]>=no2 || intervals[j][0]==no2 {
			last=intervals[j][1]
			i=j+1
			break
		}
		i=j
	}
	temp2[1]=last
	result=append(result,temp2)
	for j:=i;j<len(intervals);j++{
		result=append(result,intervals[j])
	}
	return result
}