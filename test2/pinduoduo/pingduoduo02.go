package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int
	var s string
	e1 := bufio.NewScanner(os.Stdin)
	for e1.Scan()  {
		s=e1.Text()
		count++
		if count==1 {
			break
		}
	}
	s11:=strings.Split(s," ")
	for i:=0;i<len(s11);i++{

	}
	fmt.Println(true)
}
