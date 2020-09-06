package main

import (
	"fmt"
	"index/suffixarray"
	"sort"
	"strconv"
)

func main() {
	var s="barfoothefoobarman"
	var words=[]string{"foo","bar"}
	fmt.Println(findSubstring(s,words))

}
func findSubstring(s string, words []string) []int {
	var length=len(words[0])
	var sChar=[]byte(s)
	var wordsIndex=make([][]int,len(words))
	var indexSum []int
	var indexMap []string
	var wordsIndexMap =make(map[string]int)
	index := suffixarray.New(sChar)
	for i:=0;i<len(words);i++{
		tempChar := []byte(words[i])
		offset := index.Lookup(tempChar,-1)
		for j:=0;j<len(offset);j++{
			mapIndex :=strconv.Itoa(offset[j])+strconv.Itoa(i)
			indexMap=append(indexMap,mapIndex)
			wordsIndexMap[mapIndex]=i
		}
		indexSum=append(indexSum,offset...)
		sort.Ints(offset)
		wordsIndex[i]=offset
	}
	sort.Ints(indexSum)
	var count=1
	var result []int
	for i,j:=1,1;i<=len(indexSum);{
		if count==len(words) {
			fmt.Println("count==len(words)")
			fmt.Println(j-1)
			if IsExist(j-1,len(words),wordsIndexMap,indexMap){
				result=append(result,indexSum[j])
			}
		}
		fmt.Println(indexSum[i],indexSum[i-1]+length)
		if i<len(indexSum)&&indexSum[i]==indexSum[i-1]+length {
			i++
			count++
		}else {
				count=1
				j++
				i=j
		}
	}
	fmt.Println(wordsIndex,indexSum)
	return result
}
func IsExist(j,length int,wordIndexMap map[string]int,indexMap []string) bool {
	fmt.Println(j,length,wordIndexMap,indexMap)
	var Index []int
	for i:=0;i<length ;i++  {
		Index=append(Index,wordIndexMap[indexMap[i+j]])
		fmt.Println("Index=",Index)
	}

	sort.Ints(Index)
	var i=1
	for ;i<len(Index);{
		if Index[i]==Index[i-1]+1{
			i++
		}else {
				break
		}
	}
	if i==len(Index){
		return true
	}else {
		return false
	}
}