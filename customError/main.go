package main

import "fmt"

type aError struct {
	msg  string
	code int64
}

func (e *aError) Error() string {

	return e.msg

}

func testError() error {
	return &aError{msg: "这是一个错误"}
}

func main() {

	fmt.Println(testError())

}
