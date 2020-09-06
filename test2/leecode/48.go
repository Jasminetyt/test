package main

import "fmt"

func main() {
	var matrix=[][]int{{1,2,3},{4,5,6},{7,8,9}}//{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}
	rotate(matrix)
	fmt.Println(matrix)
}
func rotate(matrix [][]int)  {
	//fmt.Println(len(matrix)/2)
	n:=len(matrix)-1
	for temp:=0;temp<len(matrix)/2;temp++{
		//fmt.Println("temp=",temp)
		for i:=temp;i<n-temp;i++{
			temp2 :=matrix[temp][i]
			matrix[temp][i]=matrix[n-i][temp]
			matrix[n-i][temp]=matrix[n-temp][n-i]
			matrix[n-temp][n-i]=matrix[i][n-temp]
			matrix[i][n-temp]=temp2
		//	fmt.Println(matrix)
		}
	}
}