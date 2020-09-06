package main

import (
	"fmt"
	"math"
)

const g_maxValue=6
func main() {
	
}
func printProbsbility(number int)  {
	if number<1{
		return
	}
	var pProbabilities =make([][]int,2)
	pProbabilities[0]=make([]int,g_maxValue*number+1)
	pProbabilities[1]=make([]int,g_maxValue*number+1)
	var flag=0
	for i:=1;i<g_maxValue ;i++  {
		pProbabilities[flag][i]=1
	}
	for k:=2;k<=number;k++{
		for i:=0;i<k;i++{
			pProbabilities[1-flag][i]=0
		}
		for i:=k;i<=g_maxValue*k;i++{
			pProbabilities[1-flag][i]=0
			for j:=1;j<=i&&j<=g_maxValue;j++{
				pProbabilities[1-flag][i]+=pProbabilities[flag][i-j]
			}
		}
		flag=1-flag
	}
	var total =math.Pow(float64(g_maxValue),float64(number))
	for i:=number;i<=g_maxValue*number;i++{
		ratio:=float64(pProbabilities[flag][i])/total
		fmt.Printf("%d:%f",i,ratio)
	}
}