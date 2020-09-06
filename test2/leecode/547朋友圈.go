package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var M=make([][]int,N)
	for i:=0;i<N;i++{
		M[i]=make([]int,N)
	}
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			fmt.Scan(&M[i][j])
		}
	}
	fmt.Println(findCircleNum(M))
}
func findCircleNum(M [][]int) int {
	var mm=make(map[int]int)
	var flag=0
	for i:=0;i<len(M);i++{
		for j:=0;j<len(M);j++{
			if M[i][j]==1{
				if _,ok:=mm[j];!ok{
					if _,ok2:=mm[i];!ok2{
					mm[i]=flag
					flag++
					}else {
						findCircleNum2(M,j,mm[i],mm)
					}

				}
			}
		}
	}
	//fmt.Println(mm)
	var max=0
	for _,v:= range mm{
		//fmt.Println(v)
		if v>max{
			max=v
		}
	}
	return max+1
}
func findCircleNum2(M [][]int,x,index int,mm map[int]int){
	for j:=0;j<len(M);j++{
		if M[x][j]==1{
			if _,ok:=mm[j];!ok{
				mm[j]=index
				findCircleNum2(M,j,index,mm)
			}
		}
	}
}