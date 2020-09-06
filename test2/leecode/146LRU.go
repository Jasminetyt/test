package main

import (
	"container/list"
	"fmt"
)

func main() {
	var cache=Constructor(10)
	cache.Put(10, 13);
	cache.Put(3, 17);
	cache.Put(6, 11);
	cache.Put(10, 5);
	cache.Put(9, 10);
	fmt.Println(cache.Get(13))
	cache.Put(2, 19);
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	cache.Put(5, 25);
	fmt.Println(cache.Get(8))
	cache.Put(9, 22);
	cache.Put(5, 5);
	cache.Put(1, 30);
	fmt.Println(cache.Get(11))
	cache.Put(9, 12);
	fmt.Println(cache.Get(7))
	fmt.Println(cache.Get(5))
	fmt.Println(cache.Get(8))
	fmt.Println(cache.Get(9))
	cache.Put(4, 30);
	cache.Put(9, 3);
	fmt.Println(cache.Get(9))
	fmt.Println(cache.Get(10))
	cache.Put(6, 14);
	cache.Put(3, 1);
	fmt.Println(cache.Get(3))
	cache.Put(10, 11);
	fmt.Println(cache.Get(8))
	cache.Put(2, 14);
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(5))
	//fmt.Println(cache.Get(1))
	//       // 返回  1
	//cache.Put(3, 3);    // 该操作会使得密钥 2 作废
	//fmt.Println(cache.Get(2))// 返回 -1 (未找到)
	//cache.Put(4, 4);    // 该操作会使得密钥 1 作废
	//fmt.Println(cache.Get(1))     // 返回 -1 (未找到)
	//fmt.Println(cache.Get(3))    // 返回  3
	//fmt.Println(cache.Get(4))



}
type LRUCache struct {
	capacity int
	nums map[int]int
	ls *list.List
	len int
}


func Constructor(capacity int) LRUCache {
	var lru LRUCache
	lru.capacity=capacity
	lru.nums=make(map[int]int)
	lru.ls=list.New()
	lru.len=0
	return lru
}


func (this *LRUCache) Get(key int) int {
	if value,ok:=this.nums[key];ok{
		this.del(key)
		return value
	}else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok:=this.nums[key];ok{
		this.nums[key]=value
		this.del(key)
	}else {
		if this.len>=this.capacity{
		temp:=this.ls.Front()
		key:=temp.Value.(int)
		delete(this.nums,key)
		this.ls.Remove(temp)
		this.len=this.len-1
	}
		this.ls.PushBack(key)
		this.len=this.len+1
		this.nums[key]=value
		this.pp()
	}


}
func (this *LRUCache) del(key int)  {
	for node:=this.ls.Front();node!=nil;node=node.Next(){
		if node.Value.(int)==key{
			this.ls.Remove(node)
			this.ls.PushBack(key)
			break
		}
	}
	this.pp()

}
func (this *LRUCache) pp( )  {
	for node:=this.ls.Front();node!=nil;node=node.Next(){
		fmt.Printf("%d ",node.Value.(int))
	}
	fmt.Println(" ")
}