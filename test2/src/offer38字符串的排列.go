package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
var str []string
scan := bufio.NewScanner(os.Stdin)
scan.Split(bufio.ScanWords)
	for scan.Scan()  {
		if scan.Text()=="end" {
			break;
		}
		str=append(str,scan.Text())
	}
permutation(0,str)
fmt.Println("")
permutation2(0,str)
fmt.Println("")
	permutation3(0,str)
	fmt.Println("")
var temp []string
for i:=1;i<=len(str);i++{
	permutation4(i,0,str,temp)
}

}
func permutation(n int,str []string)  {
	if len(str)==0 {
		return
	}
	if n==len(str)-1 {
		for i,value := range str{
			fmt.Printf("%s",value)
			if i==len(str)-1 {
				fmt.Printf("\n")
			}
		}
	}
	for i:=n;i<len(str);i++ {
		str[i],str[n]=str[n],str[i]
		permutation(n+1,str) //切片由3部分组成：指向底层数组的指针、切片访问的元素的个数（即长度）和切片允许增长的元素个数（即容量）
		str[i],str[n]=str[n],str[i] //切片内的元素发生改变后，经过函数传递，该改变也会随着函数进行传递
	}
}

func permutation2(index int,str []string)  {
	if index==len(str){
		fmt.Println(str)
		return
	}
	var mnum=make(map[string]int)
	for i:=index;i<len(str);i++{
		if _,ok:=mnum[str[i]];ok{
			continue
		}
		str[index],str[i]=str[i],str[index]
		permutation2(index+1,str)
		str[index],str[i]=str[i],str[index]
		mnum[str[i]]=1
	}
}
//排列
func permutation3(index int,str []string)  {
	var mnum=make(map[string]int)
	for i:=index;i<len(str);i++{
		if _,ok:=mnum[str[i]];ok{
			continue
		}
		str[index],str[i]=str[i],str[index]
		permutation3(index+1,str)
		str[index],str[i]=str[i],str[index]
		mnum[str[i]]=1
	}
}
//组合
func permutation4(index,x int,str []string,temp []string){
	if index==len(temp){
		fmt.Println(temp)
		return
	}
	var mnum=make(map[string]int)
	for i:=x;i<len(str);i++{
		if _,ok:=mnum[str[i]];ok{
			continue
		}
		mnum[str[i]]=1
		temp=append(temp,str[i])
		permutation4(index,i+1,str,temp)
		temp=temp[:len(temp)-1]
	}
}
