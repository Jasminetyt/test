package main

import "fmt"

func main() {
	fmt.Println(intToRoman(9))
}
func intToRoman(num int) string {
	var match=map[string]int{"I":1,"IV":4,"V":5,"IX":9,"X":10,"XL":40,"L":50,"XC":90,"C":100,"CD":400,"D":500,"CM":900,"M":1000}
	var index=[]string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	var result string
	for num>0 {
		for i:=0;i<len(index);{
			if num>=match[index[i]] {
				num=num-match[index[i]]
				result=result+index[i]
			}else {
				i++
			}
		}
	}
	return result
}