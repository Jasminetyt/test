package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	d1 := bufio.NewReader(os.Stdin)
	chars,_,_:=d1.ReadLine()
	var s =new(st)
	st := s.new1()
	var count,length=0,0
	for _,char := range chars{
		if char=='[' {
			st.push1(char)
			count=len(s.nums)
		}
		if char==']'{
			if !st.pop1(){
				return
			}
		}
		if count>length {
			length=count
		}
	}
	fmt.Println(length)
}

type st struct{
	nums []byte
}

func (s *st) new1() *st {
	s.nums=[]byte{}
	return s
}
func (s *st) pop1() bool {
	if len(s.nums) <=0{
		return false
	}
	s.nums=s.nums[0:len(s.nums)-1]
	return true
}
func (s *st) push1(char byte)  {
	s.nums=append(s.nums,char)
}
