package main

import (
	"bytes"
	"container/list"
	"encoding/binary"
	"reflect"
)

func main() {

}

func genStruct(dest interface{}) list.List {
	_list := list.New()
	_type := reflect.TypeOf(dest)
	for i := 0; i < 10; i++ {
		_list.PushBack(reflect.New(_type).Interface())
	}
	return _list
}

func FromInt(n int) []byte {
	var buf bytes.Buffer
	binary.Write(buf, binary.LitterEndian, n)
	return buf.Bytes()
}

func check(eee ...interface{}) {
	for _, e := range eee {
		if err, ok := e.(error); ok && err != nil {
			panic(err)
		}
	}
}

func noop(...interface{}) {}
