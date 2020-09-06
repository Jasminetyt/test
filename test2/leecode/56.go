package main

import (
	"fmt"
	"sort"
)

func main() {
	var intervals =[][]int{{4,5},{2,4},{4,6},{3,4},{0,0},{1,1},{3,5},{2,2}} //{1,3},{2,6},{8,10},{15,18}
	fmt.Println(merge(intervals))
}
func merge(intervals [][]int) [][]int {
	var index []int
	var numsmap=make(map[int]int)
	var j int
	for i:=0;i<len(intervals);{
		if v,ok:=numsmap[intervals[i][0]];ok{
			if v < intervals[i][1]{
				numsmap[intervals[i][0]]=intervals[i][1]
			}
			i++
		}else {
			index=append(index,intervals[i][0])
			numsmap[intervals[i][0]]=intervals[i][1]
			i++
			j++
		}

	}
	sort.Ints(index)
	//fmt.Println(index)
	//fmt.Println(numsmap)
	var count,temp1 int
	var ok bool
	for i:=1;i<len(index);i++{
		for temp1,ok=numsmap[index[count]];!ok;{
			count++
			temp1,ok=numsmap[index[count]]
		}
		temp2:=index[i]
		if temp2<=temp1{
			if numsmap[temp2]>temp1 {
				numsmap[index[count]]=numsmap[temp2]
			}
			delete(numsmap,temp2)
		}else {
			count++
		}
	}
	var result [][]int
	for i:=0;i<len(index);i++{
		if v,ok:=numsmap[index[i]];ok{
			temp :=[]int{index[i],v}
			result=append(result,temp)
		}
	}
	return result
}