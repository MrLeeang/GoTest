package main

import "fmt"

type Class struct {
	Name string
}
type User struct {
	Name string
	Age  int
	*Class
}

func (u *User) do() {
	println(u.Name)
	println("do")
	u.Name = "bbbb"
	println(u.Name)
}

func main() {
	var user1 User
	user1.Name = "fdsfsd"
	user1.do()
	println(user1.Name)
	if user1.Age != 0 {
		println(user1.Age)
	}

	user := &User{}
	user.Name = "user1"
	class := &Class{}
	user.Class = class
	user.Class.Name = "class1"
	fmt.Println(user.Class.Name)
	user2 := &User{}
	user2.Name = "user2"
	user2.Class = class
	user2.Class.Name = "class2"
	fmt.Println(user2.Class.Name)
	fmt.Println(user.Class.Name)
}
