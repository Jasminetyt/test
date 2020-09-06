package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"os"
)

func main() {
	file , err := os.Open("info.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	decoder := mahonia.NewDecoder("gbk")
	scanner := bufio.NewScanner(decoder.NewReader(file))
	var l=1
	for scanner.Scan() {
		lineText := scanner.Text()
		//if len(lineText)>66{
		//	fmt.Println(lineText,len(lineText))
		//}
		var name,ID, address1,address2,tel,str string
		nameRune := []rune(lineText)
		var y,flag,i int
		for i=0;i<len(nameRune);i++{
			if string(nameRune[i])=="	"{
				//fmt.Println("111")
				if flag==0{
					name=string(nameRune[y:i])
					//fmt.Println("name",name)
					flag++
					y=i+1
				}else if flag==1{
					ID=string(nameRune[y:i])
					flag++
					y=i+1
				}else if flag==2{
					address1=string(nameRune[y:i])
					flag++
					y=i+1
				}else if flag==3{
					address2=string(nameRune[y:i])
					flag++
					y=i+1
				}
			}
		}
		tel=string(nameRune[y:i])
		if address1==address2{
			str="兹证明"+name+"为我下属江南二所工人（身份证号码："+ID+"，联系电话："+tel+"），现住"+address2+"，需每日开展城市环卫保洁工作。"
		}else {
			str="兹证明"+name+"为我下属江南二所工人（身份证号码："+ID+"，"+address1+"，联系电话："+tel+"），现住"+address2+"，需每日开展城市环卫保洁工作。"
		}
		fmt.Println(l,str)
		l++
	}
}
