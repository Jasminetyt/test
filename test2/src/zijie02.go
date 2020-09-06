package main
import "fmt"
func main(){
	var n ,m ,c int
	fmt.Scanln(&n,&m,&c)//n代表珠子个数,m代表连续几个,c代表颜色数目
	var colormap=make(map[int][]int)
	var count,temp int
	for i:=0;i<n;i++{
		fmt.Scan(&count)
		for j:=0;j<count;j++{
			fmt.Scan(&temp)
			colormap[temp]=append(colormap[temp],i+1)
		}
	}
	fmt.Println(colormap)
	var result int
	for i:=1;i<=c;i++{
		temp2:=colormap[i]
		for j:=1;j<len(temp2);j++{
			if temp2[0]<temp2[len(temp2)-1]+m-n{
				result++
				break
			}
			if temp2[j]-temp2[j-1]<m{
				result++
				break
			}
		}
	}
	fmt.Println(result)
}