package main

import (
	"fmt"
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

type task backend.Task

func (u *task) Foo() (*http.Response, error) {
	fmt.Println("Inside foo")
	return u.Client.Get("http://google.com")
}

func (u *task) Bar() (*http.Response, error) {
	fmt.Println("Inside bar")
	return u.Client.Get("http://google.com")
}

var User task
