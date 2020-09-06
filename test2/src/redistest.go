package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err!=nil{
		fmt.Println("Connect to redis error",err)
		return
	}
	defer c.Close()
	//_,err=c.Do("set","key1","value1","EX","20")
	_,err=c.Do("hmset","key2","k1",int32(1),"k2","v2")
	v,err :=redis.Strings(c.Do("hmget","key2","k1","k2"))
	fmt.Println(v[0],v[1])
	for index,value:=range v{
		fmt.Println(index,value)
	}
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(v)
	_,err=c.Do("EXPIRE","key2","20")
}
