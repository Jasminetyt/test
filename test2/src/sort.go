package main

import (
	"fmt"
	"sort"
)

//type Interface interface {
//	// Len方法返回集合中的元素个数
//	Len() int
//	// Less方法报告索引i的元素是否比索引j的元素小
//	Less(i, j int) bool
//	// Swap方法交换索引i和j的两个元素
//	Swap(i, j int)
//}
type Person struct {
	Name string
	Age  int
}

type Persons []Person

func (p Persons) Len() int  {return len(p)}
func (p Persons) Less(i,j int) bool {return p[i].Age<p[j].Age}
func (p Persons) Swap(i, j int)  {p[i] ,p[j]=p[j],p[i]}

func main()  {
	person := [] Person{{"Lily", 20}, {"Bing", 18}, {"Tom", 23}, {"Vivy", 16}, {"John", 18}}
	sort.Sort(Persons(person))
	fmt.Println(person)
	bubbleSort(Persons(person))
	fmt.Println(person)
	num := []int{1,2,4,7,3,4}
	ints := sort.IntSlice(num) //IntSlice给[]int添加方法以满足Interface接口，以便排序为递增序列。
	insertSort(ints)
	fmt.Println(ints)
}

//https://www.jianshu.com/p/1f42f2ba6c0d

//bubbleSort 冒泡排序
func bubbleSort(data sort.Interface) {
	r := data.Len() - 1
	for i := 0; i < r; i++ {
		for j := r; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}

//insertSort 插入排序
func insertSort(data sort.Interface) {
	r := data.Len() - 1
	for i := 1; i <= r; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

//selectSort 选择排序
func selectSort(data sort.Interface) {
	r := data.Len() - 1
	for i := 0; i < r; i++ {
		min := i
		for j := i + 1; j <= r; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
