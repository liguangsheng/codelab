package main

import (
	"fmt"
	"syscall"
)

func main() {
	h, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		abort("LoadLibrary", err.Error())
	}
	defer syscall.FreeLibrary(h)
	proc, err := syscall.GetProcAddress(h, "GetVersion")
	if err != nil {
		abort("GetProcAddress", err.Error())
	}
	r, _, _ := syscall.Syscall(uintptr(proc), 0, 0, 0, 0)
	printVersion(uint32(r))

	getCurrentProcessId, err := syscall.GetProcAddress(h, "GetCurrentProcessId")
	if err != nil {
		abort("GetProcAddress", err.Error())
	}
	r2, _, _ := syscall.Syscall(uintptr(getCurrentProcessId), 0, 0, 0, 0)
	fmt.Println("CurrentProcessId: ", uint32(r2))
}

func abort(funcname string, err string) {
	panic(funcname + " failed: " + err)
}

func printVersion(v uint32) {
	major := byte(v)
	minor := uint8(v >> 8)
	build := uint16(v >> 16)
	print("windows version ", major, ".", minor, " (Build ", build, ")\n")
}
