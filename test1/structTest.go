package main

import "fmt"

type Class struct {
	Name string
}
type User struct {
	Name string
	*Class
}

func main() {
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
