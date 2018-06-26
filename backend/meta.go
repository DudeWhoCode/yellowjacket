package backend

import (
	"fmt"
	"net/http"
	"os"
	"plugin"
)

type Task struct {
	Client http.Client
}

type WebUser interface {
	Foo() (*http.Response, error)
	Bar() (*http.Response, error)
}

func LoadModule() {
	mod := "testfile.so"
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	custUser, err := plug.Lookup("User")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var user WebUser
	user, ok := custUser.(WebUser)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}
	fmt.Println("Before user.foo()")
	resFoo, _ := user.Foo()
	resBar, _ := user.Bar()
	fmt.Println("Results are : ", resFoo, resBar)

}
