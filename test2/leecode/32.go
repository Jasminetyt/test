package main
import(
	"fmt"
)

func main() {
	s:="((()()(((()"
	fmt.Println(longestValidParentheses(s))
}
func longestValidParentheses(s string) int {
	var left,right,max,temp int
	sc := []byte(s)
	for i:=0;i<len(sc);i++{
		if sc[i]=='('{
			left++
		}
		if sc[i]==')'{
			right++
		}
		if left==right{
			temp=right
		}
		if left<right{
			if max<left{
				max=left
			}
			left,right,temp=0,0,0
		}
	}
	if left>right && right-temp>max {
		max=right-temp
	}else if left>right && temp>max{
		max=temp
	}
	if left==right && left>max{
		max=left
	}
	return max*2
}
