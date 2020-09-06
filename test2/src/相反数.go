package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	var num []int
	temp:=n
	for temp>0{
		num=append(num,temp%10)
		temp=temp/10
	}
	fmt.Println(num)
	var num2=0
	flag:=true
	for i:=0;i<len(num);i++{
		if flag && num[i]==0{

		}else{
			flag=false
			num2=num2*10+num[i]
		}
	}
	fmt.Println(num2+n)
}
