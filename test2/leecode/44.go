package main

func main() {
	
}

func isMatch2(s string, p string) bool {
	if len(s)==0&&len(p)==0{
		return true
	}
	if len(s)!=0&&len(p)==0 {
		return false
	}
	if len(s)>1&&s[1]=='*'{
		if p[0]==s[0] || (p[0]=='?'&&len(s)!=0){

		}
	}
}
