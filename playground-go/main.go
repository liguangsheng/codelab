package main

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
)

type Resource struct {
	UserID int64
}

func main() {
	dsn := "aaa=bbb "
	fmt.Println(replaceDatabase(dsn, "solarland"))
}

func replaceDatabase(dsn, database string) string {
	r := regexp.MustCompile("database=\\S*")
	res := r.FindAllString(dsn, -1)
	fmt.Println(res)
	if len(res) > 0 {
		return r.ReplaceAllString(dsn, "database="+database)
	}
	return dsn + " database=" + database
}

func Append(slice *[]int) {
	*slice = append(*slice, 3)
}

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
