package main

import (
	"fmt"
	"reflect"
)

func main() {
	//for i,j :=1,1;i<10&&j<10 ; i++,j++ {
	//	fmt.Printf("i=%d,j=%d",i,j)
	//}
	var a =[...]int{1,2,3,4,5,6}
	var ints = a[1:4]


//fmt.Println(typeof(a))
//fmt.Println(typeof(ints))


fmt.Println(reflect.TypeOf(ints).String())
var b =[...]int{1,2,3}
test(&b)
fmt.Printf("111=%p ,b2=%d \n",&b,b[2])
var c=5
var p *int
p=&c
test2(p)
fmt.Printf("p=%p,*p=%d \n",&p,*p)
}
//func typeof(v interface{}) string {
//	return reflect.TypeOf(v).String()
//}


func test(b *[3]int)  {
	b[2]=4
fmt.Printf("222=%p ,b2=%d \n",&b,b[2])
}
func test2(p *int)  {
	fmt.Printf("p2=%p,*p2=%d \n",&p,*p)
	var a=7
	p=&a
	fmt.Printf("p2=%p,*p2=%d \n",&p,*p)



	}


