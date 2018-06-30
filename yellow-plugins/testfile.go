package main

import (
	"fmt"
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

type Task backend.Task

func (u *Task) Foo() (*http.Response, error) {
	fmt.Println("Inside foo")
	return u.Client.Get("http://google.com")
}

func (u *Task) Bar() (*http.Response, error) {
	fmt.Println("Inside bar")
	return u.Client.Get("http://google.com")
}

func GetTask() *Task {
	return &Task{}
}

var User Task
