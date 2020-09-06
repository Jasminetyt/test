package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch=make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	for i:=0;i<3;i++{
		go consume(ch,&wg)
	}
	go produce(ch,&wg)
	wg.Wait()
}

func produce(ch chan int,wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0;i<100;i++{
		ch<-i
	}
	close(ch)
}
func consume(ch chan int ,wg *sync.WaitGroup)  {
	defer wg.Done()
	for v:=range ch{
		fmt.Println(v)
	}

}


