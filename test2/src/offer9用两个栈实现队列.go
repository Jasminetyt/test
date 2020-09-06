package main

type ssssss struct {
	nums []interface{}
}
type qqqqq struct {
	s1 *ssssss
	s2 *ssssss
}
func main() {
	
}
func (s *ssssss) spush(num interface{})  {
	s.nums=append(s.nums,num)
}
func (s *ssssss) spop() interface{} {
	num :=s.nums[0]
	s.nums=s.nums[1:]
	return num
}
func (s *ssssss) length() int {
	return len(s.nums)
}
func (q *qqqqq) appendTail(nums []int)  {
		for i:=0;i<len(nums);i++{
			q.s1.spush(nums[i])
		}
}
func (q *qqqqq) deleteHead()  {
	if q.s1.length()==0&&q.s2.length()==0{
		return
	}
	if q.s2.length()!=0{
		q.s2.spop()
	}else {
		for q.s1.length()!=0{
			temp:=q.s1.spop()
			q.s2.spush(temp)
		}
		q.s2.spop()
	}

}