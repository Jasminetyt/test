package main
import "fmt"
func main(){
	var s string
	fmt.Scan(&s)
	var temp byte
	temp=s[0]
	var cmap=make(map [int]int)
	var count,i =1,1
	for i=1;i<len(s);i++{
		if temp==s[i]{
			count++
		}else{
			cmap[i-1]=count
			count=1
			temp=s[i]
		}
	}
	cmap[i-1]=count
	//fmt.Println(cmap)
	var sum=0
	for _,v:=range cmap{
		sum=sum+v
	}
	//fmt.Println(len(cmap),sum)
	fmt.Printf("%.2f",float32(sum)/float32(len(cmap)))
}
