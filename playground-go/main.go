package main

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type BizError struct {
	Code    int
	Message string
}

func (e *BizError) Error() string {
	return e.Message
}

func (e *BizError) BizCode() int {
	return e.Code
}

func CauseBizError(err error) *BizError {
	cause := errors.Cause(err)
	if bizerr, ok := cause.(*BizError); ok {
		return bizerr
	}
	return nil
}

type Person struct {
	Name string
	Age  int
}

func main() {
	t := time.Now().Truncate(time.Hour * 24).Add(time.Hour * 24)
	t.Date()
	fmt.Println(t)
	fmt.Println(time.Date(2020,9,32,0,0,0,0,t.Location()))
}

func foo() error {
	return errors.Wrap(bar(), "error in foo")
}

func bar() error {
	originErr := errors.New("I am the origin error")
	err := errors.WithStack(originErr)
	return errors.Wrap(err, "error in bar")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func noop(...interface{}) {}
