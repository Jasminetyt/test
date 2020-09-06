package main

import "sync"

func main() {
	
}
type singleton struct {}
var ins *singleton
var mu sync.Mutex

func getIns() *singleton {
	if ins==nil{
		mu.Lock()
		defer mu.Unlock()
		if ins==nil{
			ins=&singleton{}
		}

	}
	return ins
}

var once sync.Once
func getIns2() *singleton {
	once.Do(func() {
		ins=&singleton{}
	})
	return ins
}



