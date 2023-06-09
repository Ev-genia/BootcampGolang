package main

/*
#include <stdio.h>

int findVal(long *arr, int index) {
  return (*(arr + index));
}
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("Error: Empty slice")
	}
	if idx >= len(arr) {
		return 0, fmt.Errorf("Error: IndexOut")
	}
	if idx < 0 {
		return 0, fmt.Errorf("Error: Negativ index")
	}

	cIndex := C.int(idx)
	val := int(C.findVal((*C.long)(unsafe.Pointer(&arr[0])), cIndex))
	return val, nil
}

func main() {
	var index int
	fmt.Println("Enter index: ")
	fmt.Scan(&index)
	elem, err := getElement([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("elem[%d]: %d\n", index, elem)
}
