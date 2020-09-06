package main
import "fmt"

var sig = false
func main(){
	var N int
	fmt.Scan(&N)
	var Ai=make([]int,N)
	var temp []int
	var flag []bool
	for i:=0;i<N;i++{
		fmt.Scan(&Ai[i])
		for j:=0;j<Ai[i];j++{
			temp=append(temp,i+1)
			flag=append(flag,false)
		}
	}
//	fmt.Println(temp,flag)
	var result []int
	getresult(temp,&result,flag,-1)
	fmt.Println(result)
}
func getresult(temp []int,result *[]int,flag []bool,count int){
	if count==len(temp)-1{
		sig=true
		return
	}
	for i:=0;i<len(temp);i++{
		if !flag[i]&&(count==-1||temp[i]!=(*result)[count]){
			*result=append(*result,temp[i])
			flag[i]=true
			//fmt.Println(result,flag,count)
			getresult(temp,result,flag,count+1)
			if sig{
				return
			}
			*result=(*result)[:count+1]
			flag[i]=false
		}
	}
}
