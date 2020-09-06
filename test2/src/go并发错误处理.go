package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	var nums=[]int{3,5,2,5,6,8,1}
	var wg sync.WaitGroup
	wg.Add(len(nums))
	for i:=0;i<len(nums);i++{
	go printi(nums[i],&wg)
	}
	wg.Wait()
}
func printi(i int,wg *sync.WaitGroup)  {
	err := printi2(i)
	defer func() { //必须得写在return之前，才会在return之后执行
		wg.Done()
	}()
	if err!=nil{
		fmt.Println(err,i)
		return
	}
	fmt.Println("success",i)

}
func printi2(i int) error{
	if i<5{
		err :=errors.New("too little")
		return err
	}
	return nil
}