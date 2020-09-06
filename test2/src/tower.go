package main

import "fmt"

func main()  {
var num =[][]int{{9},{12,15},{10,6,8},{2,18,9,5},{19,7,10,4,16}}
fmt.Println(tower(num,0,0))
}
func tower(num [][]int,i ,j int) int  {
	if  i== len(num)-1{
		return num[i][j]
	}
	var max  int
	if tower(num,i+1,j)>tower(num,i+1,j+1) {
		max=tower(num,i+1,j)
	}else {
		max=tower(num,i+1,j+1)
	}
	return num[i][j]+max
}

