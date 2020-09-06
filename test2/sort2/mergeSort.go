package sort

import (
	"math"
)

type Mergesort struct {

}
func (m *Mergesort) MergeSort(nums []int)  {
	index:=float64(0)
	temp := int(math.Pow(2,index))
	for temp<len(nums){
		for i:=0;i<len(nums);i=i+2*temp{
			m.sort(nums,i,i+temp,i+2*temp) //start-mid,mid-end (不包括后者）
		}
		index++
		temp = int(math.Pow(2,index))
	}
}
func (m *Mergesort) sort(nums []int,start,mid,end int)  {
	if end>len(nums) {
		end=len(nums)
	}
	if mid>len(nums) {
		return
	}
	var temp=make([]int,end-start)
	index := 0
	var i,j int
	for i,j=start,mid;i<mid&&j<end;{
		if nums[i]<=nums[j] {
			temp[index]=nums[i]
			i++
		}else {
			temp[index]=nums[j]
			j++
		}
		index++
	}
	for i<mid{
		temp[index]=nums[i]
		i++
		index++
	}
	for j<end{
		temp[index]=nums[j]
		j++
		index++
	}
	for i=0;i<len(temp) ;i++  {
		nums[start]=temp[i]
		start++
	}
}
