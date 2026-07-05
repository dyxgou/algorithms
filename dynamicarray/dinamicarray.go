package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"log/slog"
	"unsafe"
)

const zero int = 0

type DynamicArray struct {
	addr *DynamicArray
	ptr  unsafe.Pointer
	len  int
	cap  int
}

func New(length, capability int) *DynamicArray {
	if capability <= 0 {
		panic("capability should be greater or equal to 1")
	}

	ptr := C.calloc(C.size_t(capability), C.size_t(unsafe.Sizeof(zero)))

	da := &DynamicArray{
		len: length,
		ptr: ptr,
		cap: capability,
	}

	da.addr = da
	return da
}

func (d *DynamicArray) copyCheck() {
	if d.addr != d {
		panic("ilegal use of DynamicArray: copied by value")
	}
}

func (d *DynamicArray) Get(i int) int {
	d.copyCheck()

	val := *(*int)(unsafe.Add(d.ptr, int(uintptr(i)*unsafe.Sizeof(zero))))
	slog.Info("get val", "i", i, "val", val, "newptr", int(uintptr(i)*unsafe.Sizeof(zero)))
	return val
}

func (d *DynamicArray) Set(i, val int) {
	d.copyCheck()

	ptr := (*int)(unsafe.Add(d.ptr, int(uintptr(i)*unsafe.Sizeof(zero))))

	*ptr = val
}

func (d *DynamicArray) grow() {
	if d.cap < 1024 {
		d.cap = d.cap * 2
	} else {
		// d.cap / 4
		d.cap += d.cap >> 2
	}

	d.ptr = C.reallocarray(d.ptr, C.ulong(d.cap), C.ulong(unsafe.Sizeof(zero)))
}

func (d *DynamicArray) Push(val int) {
	if d.len >= d.cap {
		d.grow()
	}

	d.Set(d.len, val)
	d.len++
}

func (d *DynamicArray) Pop() int {
	val := *(*int)(unsafe.Add(d.ptr, int(uintptr(d.len-1)*unsafe.Sizeof(zero))))
	d.len--

	return val
}

func (d *DynamicArray) Len() int {
	return d.len
}

func (d *DynamicArray) Cap() int {
	return d.cap
}
