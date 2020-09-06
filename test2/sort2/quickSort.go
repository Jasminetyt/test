package sort

type Quicksort struct {

}
func (q *Quicksort) QuickSort(nums []int,start,end int) {
	if start<end{
		index := q.sort(nums,start,end)
		q.QuickSort(nums,start,index-1)
		q.QuickSort(nums,index+1,end)
	}
}
func (q *Quicksort) sort(nums []int,start ,end int) int {
	i:=start+1
	j:=end
	for i<j{
		for i<end&&nums[i]<=nums[start]{
			i++
		}
		for start<j&&nums[j]>nums[start]{
			j--
		}
		if i<j{
			nums[i],nums[j]=nums[j],nums[i]
		}
	}
	nums[start],nums[j]=nums[j],nums[start]
	return j
}
