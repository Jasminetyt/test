                                                                                                                                                                                                                                                                                                                                                                                                                               package main

																																																																																																							   import (
																																																																																																								   "container/list"
																																																																																																								   "fmt"
																																																																																																								   "strconv"
																																																																																																							   )

func main() {
	//var nums []int
	//scan := bufio.NewScanner(os.Stdin)
	//scan.Split(bufio.ScanWords)
	//for scan.Scan(){
	//	if scan.Text()=="end" {
	//		break;
	//	}
	//	num,_:=strconv.Atoi(scan.Text())
	//	if num<0 {
	//		fmt.Printf("%d为无效数字，自动跳过该数\n",num)
	//	}else {
	//		nums=append(nums,num)
	//	}
	//}
	//if len(nums)<= 0{
	//	fmt.Println("长度太短")
	//	return
	//}
	//getTranslationCount(nums,len(nums)-1)
	//fmt.Println("")
	translate()
}
var cchar []byte
func getTranslationCount(nums []int,index int)  {
	if index==0 {
		addslice(nums[0])
		outputslice(cchar)
		deleteslice()
		return
	}
	num := nums[index]
	addslice(num)
	getTranslationCount(nums,index-1)
	num=nums[index-1]*10+num
	if num>=10&&num<=25 {
		deleteslice()
		addslice(num)
		if index>=2 {
			getTranslationCount(nums,index-2)
		}
		if index==1 {
			outputslice(cchar)
		}
		deleteslice()
	}
}
func outputslice(ch []byte)  {
	for i:=len(ch)-1;i>=0;i--{
		fmt.Printf("%c",ch[i])
	}
	fmt.Println("")
}
func addslice(num int)  {
	ch := 'a'+ num
	cchar=append(cchar,byte(ch))
}
func deleteslice()  {
	cchar=cchar[:len(cchar)-1]
}

func translate() {
	var num= 12258
	chars := []byte(strconv.Itoa(num))
	var list = list.New()
	translate2(chars,list,len(chars)-1)
}
func translate2(chars []byte,list *list.List,i int){
	if i<0{
		for node:=list.Back();node!=nil;node=node.Prev() {
			fmt.Printf("%s",node.Value.(string))
		}
		fmt.Println("")
		return
	}
	temp:=string(int(chars[i]-'0')+'a')
	list.PushBack(temp)
	translate2(chars,list,i-1)
	node:=list.Back()
	list.Remove(node)
	if i-1>=0{
		temp2:=int(chars[i-1]-'0')*10+int(chars[i]-'0')
		if temp2<=25 {
			temp=string(temp2+'a')
			list.PushBack(temp)
			translate2(chars,list,i-2)
			node=list.Back()
			list.Remove(node)
		}

	}

}