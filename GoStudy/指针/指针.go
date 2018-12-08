package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type data struct{ a int }

	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p,%v\n", p, p.a) // 直接用指针访问目标对象成员，无须转换

	//可以在 unsafe.Pointer 和任意类型指针间进行转换
	b := 0x12345678
	p1 := unsafe.Pointer(&b) // *int -> Pointer
	n := (*[4]byte)(p1)      // Pointer -> *[4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X\n", n[i])
	}

	//将 Pointer 转换成 uintptr，可变相实现指针运算
	c := struct {
		e string
		f int
	}{"aaa", 1234}
	p2 := uintptr(unsafe.Pointer(&c)) // *struct -> Pointer -> uintptr
	p2 += unsafe.Offsetof(c.f)        // uintptr + offset
	p3 := unsafe.Pointer(p2)          // uintptr -> Pointer
	px := (*int)(p3)                  // Pointer -> *int
	*px = 200                         // c.f = 200
	fmt.Printf("%#v\n", c)

}
