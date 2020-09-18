package main

import "syscall"

var kml syscall.Handle

func main() {
	h, err := syscall.LoadLibrary("kmlLIB/kmllib.dll")
	check(err)
	kml = h
	defer syscall.FreeLibrary(h)

	KeyDown("w")
}

func KeyDown(key string) int {
	proc, err := syscall.GetProcAddress(kml, "KeyDown")
	check(err)

	a1, err := syscall.UTF16PtrFromString(key)
	check(err)

	r, _, err := syscall.Syscall(proc, 1, uintptr(*a1), 0, 0)

	check(err)
	return int(r)
}

func check(v ...interface{}) {
	for _, i := range v {
		if err, ok := i.(error); ok && err != nil {
			panic(err)
		}
	}
}
