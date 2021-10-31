package main

import (
	"log"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	var str string
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	log.Printf("(main) %v, %#v", &str, str)
	log.Printf("(main) %#v", sh)

	data, len := lol()

	runtime.GC()

	sh.Data = uintptr(data)
	sh.Len = len
	log.Printf("(main) %v, %#v", &str, str)
	log.Printf("(main) %#v", sh)
}

func lol() (uint64, int) {
	var str = "hello"
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	log.Printf("lol) %v, %#v", &str, str)
	log.Printf("(lol) %#v", sh)

	return uint64(sh.Data), sh.Len
}
