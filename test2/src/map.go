package main

import (
	"fmt"
	"sort"
)

//map具有集合的属性，使用的是哈希表，必须可以比较相等的类型才能作为map的key，即
//slice、map、function这些内置类型是不能作为key的，struct类型在不包含以上字段的时候，是可以作为key的
func main() {
	m := make(map[int]string)
	m[2]="e"
	m[8]="d"
	m[0]="c"
	m[4]="b"
	m[6]="a"
	sortByKey(m)
	sortByValue(m)

}
//map具有无序性，有时候我们需要对map进行排序，对map的排序又分为对key的排序和对value的排序

//1）如果遇到的map.key是int的（或可排序的类型），直接取出map.key存到一个slice中，对slice进行sort,然后通过key-->查找到value即可
func sortByKey(m map[int]string)  {
	fmt.Println("\t排序之前")
	fmt.Printf("\tkey\tvalue\n")
	for key,value := range m{
		fmt.Printf("\t%d\t%s\n",key,value)
	}
	var keys []int
	for key := range m{
		keys=append(keys,key)
	}
	sort.Ints(keys)
	fmt.Println("\t排序之后")
	for _,key:= range keys{
		fmt.Printf("\t%d\t%s\n",key,m[key])
	}
}
//2）对value进行排序，map.key是唯一的，但value不是唯一的，存入的时候原来的数据就必定会出现覆盖的情况
//利用结构struct存入map，然后再对value进行排序，取值的时候抛弃了原map，利用struct来取值
func sortByValue(m map[int]string)  {
	fmt.Println("\t排序之前")
	fmt.Printf("\tkey\tvalue\n")
	for key,value := range m{
		fmt.Printf("\t%d\t%s\n",key,value)
	}
	type kv struct {
		key int
		value string
	}
	var temp []kv
	for key , value := range m{
		temp=append(temp,kv{key:key,value:value})
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].value<=temp[j].value //返回true，则i排在j之前
	})
	fmt.Println("\t排序之后")
	for _,v := range temp {
		fmt.Printf("\t%d\t%s\n",v.key,v.value)
	}
}