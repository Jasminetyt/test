package main

import "fmt"

func main() {
	a:= []int{1,2,3}
	b:= []string{"a","b","c"}
	c:= []interface{}{4,5,6,7,8,9}
	d := make([]interface{},len(a)+len(b))

	for i,value:=range a{
		d[i]=value
	}
	for j:=len(a);j<len(a)+len(b) ;j++  {
		d[j]=b[j-len(a)]
	}
	f:=combineArray(d,c)
	for i,value := range f{
		fmt.Println("i and value",i," ",value)
	}
}
func combineArray(a,b []interface{})map[interface{}] interface{} {
	if len(a)!=len(b){
		panic("Length is not same")
	}
	arrays := make(map[interface{}]interface{})
	for i,value := range a{
		arrays[value]=b[i]
	}
	return arrays
}