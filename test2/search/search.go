package search

//二分查找
func SearchK(nums []int,start,end,k int) (index int) {
	mid := (start+end)/2
	index=-1
	for start<end{
		if nums[mid]==k {
			index=mid
			break
		}
		if nums[mid]<k{
			start=mid+1
		}
		if nums[mid]>k{
			end=mid-1
		}
		mid=(start+end)/2
	}
	if nums[mid]==k{
		index=mid
	}

	return index
}
