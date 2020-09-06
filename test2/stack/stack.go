package stack

import (
	"sync"
)


type Stack struct {
	items []interface{} //切片
	lock sync.Mutex
	Stack1 *Stack
	Stack2 *Stack
}

//初始化栈
func (stack *Stack)New() *Stack {
	stack.items=[]interface{}{}
	stack.Stack1=new(Stack)
	stack.Stack2=new(Stack)
	return stack
}
//入栈
func (stack *Stack) Push(item interface{})  {
	if stack==nil {
		return
	}
	stack.lock.Lock()
	stack.items=append(stack.items, item)
	stack.lock.Unlock()
}
//出栈
func (stack *Stack) Pop() interface{} {
	if stack.Length()==0 {
		return nil
	}
	stack.lock.Lock()
	item := stack.items[len(stack.items)-1]
	stack.items=stack.items[:len(stack.items)-1]
	stack.lock.Unlock()
	return item
}
//判断栈是否为空
func (stack *Stack) IsEmpty() bool  {
	if stack==nil {
		stack=new(Stack)
		return true
	}
	return len(stack.items)==0
}
//获取栈的第一个元素，不移除
func (stack *Stack) Front() interface{} {
	if stack.Length()==0 {
		return nil
	}
	stack.lock.Lock()
	item := stack.items[len(stack.items)-1]
	stack.lock.Unlock()
	return item
}
//栈长度
func (stack *Stack) Length() int  {
	return len(stack.items)
}
//清空栈
func (stack *Stack) Clear()  {
	stack.lock.Lock()
	stack.items=stack.items[0:0]
	stack.lock.Unlock()
}
//获得栈中元素的最小值
func (stack *Stack) StackWithMin() interface{} {
	return stack.Stack2.Front()
}
func (stack *Stack) StackWithMinPush(item interface{}){
	stack.Stack1.Push(item)
	if stack.Stack2.IsEmpty() || item.(int)<stack.Stack2.Front().(int){
		stack.Stack2.Push(item)
	}else {
		stack.Stack2.Push(stack.Stack2.Front())
	}
}
func (stack *Stack) StackWithMinPop()  {
	stack.Stack1.Pop()
	stack.Stack2.Pop()
}
//遍历栈（从栈底到栈顶）
func (stack *Stack) TraceStack() []interface{} {
	if stack==nil {
		return nil
	}else {
		return stack.items
	}
}


