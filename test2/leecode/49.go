package main

import (
	"fmt"
	"sort"
)

func main() {
	var strs=[]string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

	//var s=[]byte{'w','e','l'}
	//sort.Sort(num(s))
	//fmt.Println(string(s))
}

type num []byte

func (n num) Len() int {return len(n)}
func (n num) Less(i,j int) bool {return n[i]<n[j]}
func (n num) Swap(i,j int)  {n[i],n[j]=n[j],n[i]}
func groupAnagrams(strs []string) [][]string {
	var mstr=make(map[string][]int)
	for i:=0;i<len(strs);i++{
		temp := []byte(strs[i])
	    sort.Sort(num(temp)) //将每个字符串进行按字母排序，字母异位词经排序后是相同的
		mstr[string(temp)]=append(mstr[string(temp)],i) //将字母异位词在字符串数组中的下标存入map中
	}

	var result=make([][]string,len(mstr))
	   i:=0
		for _,temp2:= range mstr{
			for j:=0;j<len(temp2);j++{
				result[i]=append(result[i],strs[temp2[j]])
			}
			i++
		}

	return result
}


