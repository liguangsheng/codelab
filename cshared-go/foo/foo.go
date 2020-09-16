package foo

import (
	"C"
	"fmt"
)

//export foo
func foo() {
	fmt.Println("this is func foo")
}
