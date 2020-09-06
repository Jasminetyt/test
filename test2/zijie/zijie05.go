package main

import "fmt"

func main() {
	var board=[][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	var flag=make([][]bool,3)
	for i:=0;i<len(flag);i++{
		flag[i]=make([]bool,4)
	}
	var s string
	fmt.Scan(&s)
	ss:=[]byte(s)

	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			if board[i][j]==ss[0]{
				flag[i][j]=true
				if lookfor(board,i,j,1,ss,flag){
					fmt.Println(true)
					return
				}else {
					flag[i][j]=false
				}
			}
		}
	}
	fmt.Println(false)
}
func lookfor(board [][]byte,x,y,index int,ss []byte,flag [][]bool) bool {
	//fmt.Println(x,y,index)
	if index==len(ss){
		return true
	}
	if y+1<len(board[0])&&board[x][y+1]==ss[index]&&!flag[x][y+1]{
		flag[x][y+1]=true
		if lookfor(board,x,y+1,index+1,ss,flag){
			return true
		}else {
			flag[x][y+1]=false
		}
	}
	if y-1>=0&&board[x][y-1]==ss[index]&&!flag[x][y-1]{
		flag[x][y-1]=true
		if lookfor(board,x,y-1,index+1,ss,flag){
			return true
		}else {
			flag[x][y-1]=false
		}
	}
	if x+1<len(board)&&board[x+1][y]==ss[index]&&!flag[x+1][y]{
		flag[x+1][y]=true
		if lookfor(board,x+1,y,index+1,ss,flag){
			return true
		}else {
			flag[x+1][y]=false
		}
	}
	if x-1>=0&&board[x-1][y]==ss[index]&&!flag[x-1][y]{
		flag[x-1][y]=true
		if lookfor(board,x-1,y,index+1,ss,flag){
			return true
		}else {
			flag[x-1][y]=false
		}
	}
	return false
}
