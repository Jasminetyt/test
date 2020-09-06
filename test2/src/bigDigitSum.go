package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a,b string
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//fmt.Println(input.Text())
	fmt.Scanf("%s",&a)
	fmt.Scanf("%s",&b)
	//fmt.Println(a+","+b)
	chara := []byte(a)
	charb := []byte(b)
	//fmt.Printf("%c,%c\n",chara,charb)
	var i,j = len(chara)-1,len(charb)-1
	var num1,num2 int
	var flag,num=0,0
//	var result bytes.Buffer
	var str=""
	for i>=0&&j>=0 {
		num1=int(chara[i]-'0')
		num2=int(charb[j]-'0')
		if (num2+num1+flag)>9 {
			num=(num1+num2+flag)-10
			flag=1
		}else {
			num=num1+num2+flag
			flag=0
		}
		//result.WriteString(string(num))
		str=strconv.Itoa(num)+str
		i--
		j--
	}
	if i>=0 {
		num1=int(chara[i]-'0')
		num=num1+flag
		flag=0
		//result.WriteString(string(num))
		str=strconv.Itoa(num)+str
		i--
	}
	if j>=0 {
		num2=int(charb[j]-'0')
		num=num2+flag
		flag=0
		//result.WriteString(string(num))
		str=strconv.Itoa(num)+str
		j--
	}
//	fmt.Println(result)
	fmt.Println(str)
}
