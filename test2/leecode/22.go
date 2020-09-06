package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(1))
}
func generateParenthesis(n int) []string {
	if n<=0 {
		return nil
	}
	var temp=make([]byte,n*2)
	var result []string
	printString(0,2*n,n,n,temp,&result)
	return result
}
func printString(count,sum,left,right int,temp []byte,result *[]string)  {
	if count==sum {
		*result=append(*result,string(temp))
	}else {
		if left<0 || right<0 || left>right {
			return
		}
		if left>0 {
			temp[count]='('
			printString(count+1,sum,left-1,right,temp,result)
		}
		if right>0 {
			temp[count]=')'
			printString(count+1,sum,left,right-1,temp,result)
		}
	}
}
