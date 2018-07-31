package main

import (
	"fmt"
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

type Task backend.User

func (u *Task) Foo() (*http.Response, error) {
	fmt.Println("Inside foo")
	return u.Client.Get("https://www.twitter.com/dudewhocode")
}

func (u *Task) Bar() (*http.Response, error) {
	fmt.Println("Inside bar")
	return u.Client.Get("https://www.twitter.com/spf13")
}

func (u *Task) Baz() (*http.Response, error) {
	fmt.Println("Inside baz")
	return u.Client.Get("https://www.twitter.com/cnu")
}

func GetTask() *Task {
	return &Task{}
}

var User Task
