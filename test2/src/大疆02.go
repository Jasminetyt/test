package main

import (
	"fmt"
	"sort"
)

func main() {
	var N,T int//种类，钱
	var V,I,C int//价格，满意度，数量
	var result []int

		 fmt.Scan(&N,&T)

			var num=make([][]int,N)
			var mnum=make(map[float64][]int)
			var index=make([]float64,N)
			for i:=0;i<N;i++{
				fmt.Scan(&V,&I,&C)
				num[i]=make([]int,3)
				num[i][0]=V
				num[i][1]=I
				num[i][2]=C
				if value,ok:=mnum[float64(I)/float64(V)];ok{
					value=append(value,i)
					mnum[float64(I)/float64(V)]=value
				}else {
					var value []int
					value=append(value,i)
					mnum[float64(I)/float64(V)]=value
				}

				index[i]=float64(I)/float64(V)
			}
			sort.Float64s(index)
			fmt.Println(index)
			var temp2 ,sum int
			for i:=N-1;i>=0&&T>=0;i--{
				temp:=mnum[index[i]]
				T,temp2=getcount(T,num,temp)
				fmt.Println(temp2)
				sum=sum+temp2
			}
			result=append(result,sum)


	for i:=0;i<len(result);i++{
		fmt.Println(result[i])
	}
}
func getcount(T int,num [][]int,temp []int) (int,int) {
	var m int
	var index []int
	var mm=make(map[int][]int)
	for i:=0;i<len(temp);i++{
		index=append(index,num[temp[i]][0])
		if value,ok:=mm[num[temp[i]][0]];ok{
			value=append(value,i)
			mm[num[temp[i]][0]]=value
		}else {
			var value []int
			value=append(value,i)
			mm[num[temp[i]][0]]=value
		}
	}
	sort.Ints(index)
	fmt.Println(index)
	for i:=0;i<len(index);i++{
		temp:=mm[index[i]]
		m=0
		for j:=0;j<len(temp);j++{
			for z:=0;z<num[temp[j]][2]&&T>0;z++{
				T=T-num[temp[j]][1]
				fmt.Println(num[temp[j]][1])
				m=m+num[temp[j]][1]
			}
		}
	}
	return T,m
}
