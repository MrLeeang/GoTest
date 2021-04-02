package main

import (
	"fmt"
)

type A interface {
	testA()
}
type B interface {
	testB()
}
type C interface {
	testC()
	A
	B
}

type Test struct {
}

func (t *Test) testA() {
	fmt.Println("testA")

}

func (t *Test) testB() {
	fmt.Println("testB")

}

func (t *Test) testC() {
	fmt.Println("testC")

}

func createTest() C {
	return &Test{}
}

func main() {
	t := &Test{}
	t.testA()
	t.testB()
	t.testC()

	// 申明接口A
	var a1 A = t
	a1.testA()
	// 申明接口B
	var b1 B = t
	b1.testB()
	// 申明接口C
	var c1 C = t
	c1.testA()
	c1.testB()
	c1.testC()

	// 通过返回值来指定接口类型
	tt := createTest()
	tt.testA()
	tt.testB()
	tt.testC()
}
