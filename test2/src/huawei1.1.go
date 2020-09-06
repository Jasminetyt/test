package main

import "fmt"

func main() {
	var n int

	for {
		a, _ := fmt.Scan(&n)
		if a == 0 {
			break
		} else {
			if n<1 {
				return
			}
			if n>1000 {
				n=1000
			}
			var last=0
			for i:=2;i<=n;i++{
				last=(last+3)%i
			}
			fmt.Println(last)
		}
	}
	}




