package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("第1种输入：")
    var a1,a2 int
	fmt.Scan(&a1,&a2)//当输入的数据与参数相符时输入结束
	fmt.Println(a1)//通过空格来分词，直到接受完所有指定的变量数，Scan函数才会返回，回车符也无法提前让它返回

	fmt.Println("第2种输入：")
    var b1,b2 string
    fmt.Scanln(&b1,&b2)//以换行作为结束标志，空格作为分隔符
    fmt.Println(b1,b2)//\n会让函数提前返回，将返回时还未接收到的值的变量赋为空

	fmt.Println("第3种输入：")
	var c1,c2,c3,c4  string
	fmt.Scanf("%s,%s",&c1,&c2)//输入需与定义的格式相同
	fmt.Println(c1,c2)
	fmt.Scanf("%s;%s",&c3,&c4)//需要指定输入的格式，适用于完全了解输入格式的场景，可以直接把不需要的部分过滤掉
	fmt.Println(c1,c2,c3,c4)//如果输入不符合指定的格式，从不符合处开始，其后的变量值为空

	fmt.Println("第4种输入：")
	d1 := bufio.NewReader(os.Stdin)
	input ,_ := d1.ReadString('S')//分隔符为字节类型
	fmt.Println(input)

	fmt.Println("第5种输入：")
	e1 := bufio.NewScanner(os.Stdin)
	for e1.Scan()  {
		fmt.Println(e1.Text())//以换行作为一次的输入结束，该程序将无限输入
	}

	 fmt.Println("第6种输入：")
	 var f1="123 / ert / 5.2"
	 var formate = "%d / %s /%f"
	 var(
	 	f2 int
	 	f3 string
	 	f4 float32
	 )
	 fmt.Sscanf(f1,formate,&f2,&f3,&f4)//从变量中输入
	 fmt.Println(f2,f3,f4)
}
