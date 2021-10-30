package main

import (
	"log"
	"reflect"
	"unsafe"
)

func main() {
	var slice []string
	addString(&slice)
	log.Printf("--- from main ===")
	for _, str := range slice {
		log.Printf("%v, %v", &str, str)
		sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
		log.Printf("%#v", sh)
	}
}

func addString(slice *[]string) {
	var str = "hello"
	log.Printf("%v, %v", &str, str)
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	log.Printf("%#v", sh)
	*slice = append(*slice, str)
}
