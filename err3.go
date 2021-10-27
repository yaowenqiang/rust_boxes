package main

import (
	"log"
)

func main() {
	var err error
	err = (*naughtyError)(nil)
	log.Printf("%v", err)

	err = (*niceError)(nil)
	log.Printf("%v", err)
}

type naughtyError struct{}

func (ne *naughtyError) Error() string {
	return "oh no"
}

type niceError struct{}

func (ne *niceError) Error() string {
	return "ho ho ho!"
}
