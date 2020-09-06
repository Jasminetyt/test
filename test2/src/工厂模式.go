package main

import "fmt"

type operation interface {
	doOperation(num1,num2 int) int
}
type sumOperation struct {}
type subOperation struct {}
type mulOperation struct {}
type factory struct {
	opera operation
}
func (sum *sumOperation) doOperation(num1,num2 int) int{
	return num1+num2
}
func (sum *subOperation) doOperation(num1,num2 int) int{
	return num1-num2
}
func (sum *mulOperation) doOperation(num1,num2 int) int{
	return num1*num2
}
func (fac *factory) getfactoeyresult(operationtype string) operation {
	if operationtype=="sum"{
		return new(sumOperation)
	}else if operationtype=="sub"{
		return new(subOperation)
	}else if operationtype=="mul"{
		return new(mulOperation)
	}else {
		return nil
	}
}
func main() {
	var fac=new(factory)
	op1:=fac.getfactoeyresult("sum")
	fmt.Println(op1.doOperation(1,2))
	op2:=fac.getfactoeyresult("sub")
	fmt.Println(op2.doOperation(1,2))
	op3:=fac.getfactoeyresult("mul")
	fmt.Println(op3.doOperation(1,2))
}
