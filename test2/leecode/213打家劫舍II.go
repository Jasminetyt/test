package main

func main() {
	//相邻会报警，房屋排列为圆圈，首位相接，求最大价值
}
func rob2(nums []int) int {
	if nums==nil || len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	if len(nums)==2{
		if nums[0]>nums[1]{
			return nums[0]
		}else{
			return nums[1]
		}
	}
	var dp1=make([]int,len(nums)-1)
	dp1[0]=nums[0]
	if nums[0]>nums[1]{
		dp1[1]=nums[0]
	}else{
		dp1[1]=nums[1]
	}


	for i:=2;i<len(nums)-1;i++{
		if dp1[i-2]+nums[i]>dp1[i-1]{
			dp1[i]=dp1[i-2]+nums[i]
		}else{
			dp1[i]=dp1[i-1]
		}
	}
	// fmt.Println(dp1)
	var dp2=make([]int,len(nums)-1)
	dp2[0]=nums[1]
	if nums[1]>nums[2]{
		dp2[1]=nums[1]
	}else{
		dp2[1]=nums[2]
	}
	for i:=2;i<len(nums)-1;i++{
		if dp2[i-2]+nums[i+1]>dp2[i-1]{
			dp2[i]=dp2[i-2]+nums[i+1]
		}else{
			dp2[i]=dp2[i-1]
		}
	}
	// fmt.Println(dp2)
	if dp1[len(nums)-2]>dp2[len(nums)-2]{
		return dp1[len(nums)-2]
	}else{
		return dp2[len(nums)-2]
	}

}