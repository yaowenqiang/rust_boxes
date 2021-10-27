package main

import (
	"log"
	"unsafe"
)

func readIssue() (string, error) {
	var err *naughtyError
	log.Printf("(in readIssue) nil? %v, size = %v", err == nil, unsafe.Sizeof(err))
	return "", err
}

func main() {
	issue, err := readIssue()
	log.Printf("(in main) nil? %v, size = %v", err == nil, unsafe.Sizeof(err))

	if err != nil {
		log.Fatalf("fatal error: %+v", err)
	}
	log.Printf("issue = %v", issue)
}

type naughtyError struct{}

func (ne *naughtyError) Error() string {
	return "oh no"
}
