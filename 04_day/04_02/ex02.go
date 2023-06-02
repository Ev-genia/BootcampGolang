package main

// #include "cow.c"
import "C"
import (
	"fmt"
	"unsafe"
)

// import "unsafe"

func main() {
	cs := C.CString("test")
	buff := C.ask_cow(cs)
	fmt.Println("buff: ", buff)

	defer C.free(unsafe.Pointer(cs))
}
