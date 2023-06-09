package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#include "window.h"
#include "application.h"

*/
import "C"
import "unsafe"

func main() {
	cs := C.CString("School 21")
	defer C.free(unsafe.Pointer(cs))
	wind := C.Window_Create(C.int(320), C.int(250), C.int(500), C.int(300), cs)
	defer C.free(unsafe.Pointer(wind))

	C.InitApplication()
	C.Window_MakeKeyAndOrderFront(wind)
	C.RunApplication()
}
