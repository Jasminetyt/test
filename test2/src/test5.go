package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	t=time.Now()
	var result=t.Unix()
	fmt.Println(result,int64(t.Nanosecond())/int64(time.Microsecond))
}
