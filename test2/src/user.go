package main

import "fmt"

type notifer interface {
	notify()
}
type user struct {
	name string
	eamil string
}
type admin struct {
	user
	level string
}
func (u user) notify()  {
	fmt.Printf("Sending user eamil to %s<%s>\n",u.name,u.eamil)
}
func (a admin) notify()  {
	fmt.Printf("Sending admin eamil to %s<%s>\n",a.name,a.eamil)
}
func sendNotification(n notifer)  {
	n.notify()
}
func main() {
	user:=user{
		name:"ttt",
		eamil:"ttt@163.com",
	}
	ad := &admin{
		user:user,
		level:"super",
	}
	ad.user.notify()
	sendNotification(ad.user)
}

