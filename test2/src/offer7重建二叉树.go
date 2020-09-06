package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"test2/tree"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var str=make([]string,2)
	var i int
	for scan.Scan(){
		str[i]=scan.Text()
		i++
		if i==2{
			break
		}
	}
	str1:=strings.Split(str[0]," ")
	str2:=strings.Split(str[1]," ")
	var num1,num2 []interface{}
	for i:=0;i<len(str1);i++{
		temp,_:=strconv.Atoi(str1[i])
		num1=append(num1,temp)
	}
	for i:=0;i<len(str2);i++{
		temp,_:=strconv.Atoi(str2[i])
		num2=append(num2,temp)
	}
	var tr *tree.Tree
	root := tr.CreateTreeByPreOrderAndInOrder(num1,num2,0,len(num1)-1,0,len(num2)-1)
	tr.PrintTreeByLine(root)
}
type offer7tree struct {
	value interface{}
	left *offer7tree
	right *offer7tree
}

func createtree(pre,in []int,pstart,pend,istart,iend int) *offer7tree {
	var node *offer7tree
	if pstart<pend{
		node=new(offer7tree)
		node.value=pre[pstart]
		index:=searchi(pre[pstart],in,istart,iend)
		node.left=createtree(pre,in,pstart+1,pstart+index-istart+1,istart,index-1)
		node.right=createtree(pre,in,pstart+index-istart+2,pend,index+1,iend)
	}
	return node
}

func searchi(num int,in []int,istart,iend int)int{
	var i int
	for i=istart;i<=iend;i++{
		if in[i]==num{
			break
		}
	}
	return i
}

