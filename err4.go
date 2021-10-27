package main

import (
	"errors"
	"log"
)

func showType(err error) {
	if _, ok := err.(*naughtyError); ok {
		log.Printf("got a *naughtyError")
	} else if _, ok := err.(*niceError); ok {
		log.Printf("got a *niceError")
	} else {
		log.Printf("got another kind of error")
	}

}

func main() {
	showType((*naughtyError)(nil))
	showType((*niceError)(nil))
	showType(errors.New(""))
}

type naughtyError struct{}

func (ne *naughtyError) Error() string {
	return "oh no"
}

type niceError struct{}

func (ne *niceError) Error() string {
	return "ho ho ho!"
}
