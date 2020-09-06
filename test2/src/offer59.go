package main

import (
	"bufio"
	"fmt"
	"os"
	"test2/queue"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	chars ,_,_ := read.ReadLine()
	var nums []int
	var windows int
	for _,value:=range chars{
		nums=append(nums,int(value)-'0')
	}
	if len(nums)<=0{
		fmt.Println("数组为空")
		return
	}
	fmt.Printf("输入窗口大小：")
	fmt.Scan(&windows)
	if windows>len(nums) {
		fmt.Println("窗口长度大于数组长度")
		return
	}
	if windows==0 {
		fmt.Println("窗口长度为0")
		return
	}
	maxInWindows(nums,windows)
}
func maxInWindows(nums []int,window int)  {
	var q =new(queue.Queue)
	q=q.New()
	var count int
	for i:=0;i<len(nums);i++{
		count++
		if count>window {//插入元素需大于窗口长度才能有最大值
			fmt.Printf("%d ",nums[q.Front().(int)])
		}
		if !q.IsEmpty() {
			temp := q.Front().(int)
			if i-temp>=window { //队列首元素于先元素长度大于窗口长度时，移除首元素
				q.DeQueue()
			}
			if !q.IsEmpty() {
				if temp = q.Front().(int);nums[temp]<=nums[i]{ //若插入元素大于队列首元素，则清空队列
					q.Clear()
				}
			}
			for  !q.IsEmpty() {
				temp =q.Last().(int)
				if nums[i]>nums[temp] { //若当前元素小于队列首元素但是大于队列中的一部分元素，则将这些元素移除
					q.Remove()
				}else {
					break
				}
			}
		}

		q.EnQueue(i)
	}
	fmt.Println(nums[q.Front().(int)])
}
func queueWithMax(nums []int )  {
	var q=new(queue.Queue)
	q.New()
	for i:=0;i<len(nums);i++{
		if !q.IsEmpty() {

		}
	}
}
