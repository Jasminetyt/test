package main

import "fmt"

type addoperationi interface {
	doAddOperatoin(num1,num2 int) int
}
type suboperationi interface {
	doSubOperatoin(num1,num2 int) int
}
type addoperation struct {}
type suboperation struct {}
type addchangesub struct {
	sub suboperationi
}
func (add *addoperation) doAddOperatoin(num1,num2 int) int {
	return num1+num2
}
func (add *suboperation) doSubOperatoin(num1,num2 int) int {
	return num1-num2
}
func (acs *addchangesub) doAddOperatoin(num1,num2 int) int {
	temp:=&addchangesub{new(suboperation)}
	return temp.sub.doSubOperatoin(num1,num2)
}
func main() {
	var oper1 =new(addoperation)
	fmt.Println(oper1.doAddOperatoin(1,2))
	var oper2=new(suboperation)
	fmt.Println(oper2.doSubOperatoin(1,2))
	var oper3=new(addchangesub)
	fmt.Println(oper3.doAddOperatoin(1,2))
}
