package queue

import (
	"sync"
)


type Queue struct {
	items []interface{}
	lock sync.Mutex
}

func (queue *Queue) New() *Queue {
	queue.items=[]interface{}{}
	return queue
}
//入队列
func (queue *Queue) EnQueue(item interface{})  {
	queue.lock.Lock()
	queue.items=append(queue.items,item)
	queue.lock.Unlock()
}
//出队列
func (queue *Queue) DeQueue() interface{}  {
	queue.lock.Lock()
	item := queue.items[0]
	queue.items=queue.items[1:]
	queue.lock.Unlock()
	return item
}
//获取队列的第一个元素，不移除
func (queue *Queue) Front() interface{} {
	queue.lock.Lock()
	item := queue.items[0]
	queue.lock.Unlock()
	return item
}
//判断队列是否为空
func (queue *Queue) IsEmpty() bool {
	return len(queue.items)==0
}
//获取队列长度
func (queue *Queue) Size() int {
	return len(queue.items)
}
//清空队列
func (queue *Queue) Clear()  {
	queue.lock.Lock()
	queue.items=queue.items[0:0]
	queue.lock.Unlock()
}
//获取队列最后一个元素，不移除
func (queue *Queue) Last() interface{} {
	queue.lock.Lock()
	item := queue.items[len(queue.items)-1]
	queue.lock.Unlock()
	return item
}
func (queue *Queue) Remove()  {
	queue.lock.Lock()
	queue.items=queue.items[0:len(queue.items)-1]
	queue.lock.Unlock()
}