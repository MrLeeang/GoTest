package main

import "fmt"

type AInterface interface {
	a(n string) Binterface
}

type Binterface interface {
	c() Cinterface
}

type Cinterface interface {
	list()
}

type aa struct {
}

type cc struct {
}

type bb struct {
}

func (bb *bb) c() Cinterface {
	return &cc{}
}

func (cc *cc) list() {
	fmt.Print("cc list")
}

func (aa *aa) a(n string) Binterface {
	fmt.Print(n)
	return &bb{}
}

func main() {
	aa1 := &aa{}
	aa1.a("fds").c().list()
}
