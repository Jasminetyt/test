package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	var num=make([]int,n)
	var result=make([]string,n)
	for i:=0;i<n;i++{
		fmt.Scan(&num[i])
		result[i]=getresult2(num[i])
	}
	for i:=0;i<n;i++{
		fmt.Println(result[i])
	}
}
func getresult2(num int) string{
	mnum,count := getmap(num)
	for i:=2;i<10;i++{
		tmap,count2 := getmap(num*i)
		if count2>count{
			return "Impossible"
		}
		if compare(mnum,tmap){
			return "Possible"
		}
	}
	return "Impossible"
}
func getmap(num int) (map[int]int, int){
	temp:=num
	var count int
	var mnum=make(map [int]int)
	for temp>0{
		if v,ok:=mnum[temp%10];ok{
			mnum[temp%10]=v+1
		}else{
			mnum[temp%10]=1
		}
		temp=temp/10
		count++
	}
//	fmt.Println(num,mnum)
	return mnum,count
}
func compare(mnum,tmap map[int]int) bool {
	for index,v:=range mnum{
		if v!=tmap[index]{
			return false
		}
	}
	return true
}
