package main

import "fmt"

type strategy interface {
	doOperation(num1,num2 int) int
}
type operationAdd struct {}
type operationSub struct {}
type operationMul struct {}
type context struct {
	operation strategy
}

func (add *operationAdd) doOperation(num1,num2 int) int {
	return num1+num2
}
func (add *operationSub) doOperation(num1,num2 int) int {
	return num1-num2
}
func (add *operationMul) doOperation(num1,num2 int) int {
	return num1*num2
}
func (cont *context) getoperationresult(strate strategy,num1,num2 int) int {
	st:=&context{operation:strate}
	return st.operation.doOperation(num1,num2)
}
func main() {
	var con=new(context)
	fmt.Println(con.getoperationresult(new(operationAdd),1,2))
	fmt.Println(con.getoperationresult(new(operationSub),1,2))
	fmt.Println(con.getoperationresult(new(operationMul),1,2))

}
