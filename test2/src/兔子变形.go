package main

import (
	"fmt"
)

//有一对兔子，7个月后成熟生一对兔子，之后每3个月生一对，问n个月后有几对兔子
func main() {
	var month int
	fmt.Scan(&month)
	var count,index=0,1
	var parent=make(map[int]int)
	var child=make(map[int]int)
	child[0]=count
	for count<=month{
		for key,value:=range child{
			value=value+1
			if value==7 {
				delete(child,key)
				parent[key]=-1
			}else {
				child[key]=value
			}
		}
		//fmt.Println(parent)
		for key,value:=range parent{
			value=value+1
			if value==3{
				child[index]=0
				index++
				parent[key]=0
			}else if value==0{
				child[index]=0
				index++
				parent[key]=value
			}else{
				parent[key]=value
			}
		}
		count++
	}
	fmt.Println(child,parent)
	fmt.Println(len(child)+len(parent))
}
