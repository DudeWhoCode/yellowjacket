package main

import (
	"fmt"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

type task backend.Task

func (u *task) Foo() {
	fmt.Println("Inside foo")
	u.Client.Get("http://google.com")
}

func (u *task) Bar() {
	fmt.Println("Inside bar")
	u.Client.Get("http://google.com")
}

var User task
