package main

import "fmt"

func main() {
	var obstacleGrid=[][]int{{0},{0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var temp=make([]int,len(obstacleGrid[0]))
	var value=1
	for i:=0;i<len(obstacleGrid[0]);i++{
		if obstacleGrid[0][i]==1{
			value=0
		}
		temp[i]=value
	}
	var left int
	for i:=1;i<len(obstacleGrid);i++{
		for j:=0;j<len(obstacleGrid[0]);j++{
			if obstacleGrid[i][j]==1{
				temp[j]=0
			}else {
				if j-1<0{
					left=0
				}else {
					left=temp[j-1]
				}
				temp[j]=temp[j]+left
			}
		}
	}
	return temp[len(obstacleGrid[0])-1]
}