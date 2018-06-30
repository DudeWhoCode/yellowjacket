package backend

import (
	"fmt"
	"net/http"
	"os"
	"plugin"
	"reflect"
)

type Task struct {
	Client http.Client
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
		fmt.Println("Lookup error: ", err)
		os.Exit(1)
	}
	fooType := reflect.TypeOf(custUser)
	fooVal := reflect.ValueOf(custUser)
	var methods []string
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		methods = append(methods, method.Name)
	}
	for _, m := range methods {
		res := fooVal.MethodByName(m).Call([]reflect.Value{})
		response := res[0].Interface().(*http.Response)
		fmt.Println(response.StatusCode)
	}
}
