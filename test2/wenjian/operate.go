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
	//line := bufio.NewReader(decoder.NewReader(file))
	////line := bufio.NewReader(file)
	//
	//for {
	//
	//	content , _ , err := line.ReadLine()
	//	if err == io.EOF{
	//		break
	//	}
	//	//for i:=0;i< len(content);i++{
	//	//	fmt.Print(i,string(content))
	//	//}
	//
	//	fmt.Println( string(content))
	//	ba := []byte{}
	//	for _,b:= range content{
	//		ba=append(ba,byte(b))
	//	}
	//	if len(ba)>66{
	//
	//
	//		fmt.Println("mmmm",string(ba),len(ba))
	//	}
	//
	//	//fmt.Println(reflect.TypeOf(content))
	//
	//
	//}
	scanner := bufio.NewScanner(decoder.NewReader(file))
	for scanner.Scan(){
		lineText := scanner.Text()
		//if len(lineText)>66{
		//	fmt.Println(lineText,len(lineText))
		//}
		nameRune := []rune(lineText)
		var x,y int
		if len(nameRune)>22{
		//fmt.Println(len(nameRune),string(nameRune))
	roof:for i:=0;i<len(nameRune);i++{
		if string(nameRune[i])=="现"&&string(nameRune[i+1])=="住"{
			x=i+2
			for j:=i+2;j<len(nameRune);j++{
				if string(nameRune[j])=="，"{
					y=j
					break roof
				}
			}
		}
	}


			fmt.Println(string(nameRune[x:y]))
		//	fmt.Println(string(nameRune))


	}
	}
}
