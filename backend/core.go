package backend

import (
	"fmt"
	"net/http"
	"os"
	"plugin"
	"reflect"
	"sync"
)

type swarm struct {
	WaspCount int
	HatchRate int
	// Inputs  SwarmInput
	OutChan chan RawResponse
	Wasp    *UserBehaviour
	NumReq  int
	NumFail int
}

type User struct {
	Client http.Client
}

type UserBehaviour struct {
	Behaviour plugin.Symbol
	Tasks     []string
}

var s *Swarm

func GetSwarm(input SwarmInput) (s *Swarm) {
	if s == nil {
		s = &Swarm{
			input,
			make(chan RawResponse),
			&UserBehaviour{
				nil,
				nil,
			},
			0,
			0,
		}
	}
	return
}

var Wasp *UserBehaviour

func (u *UserBehaviour) GetMethods() (methods []string) {
	fooType := reflect.TypeOf(u.Behaviour)
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		methods = append(methods, method.Name)
	}
	u.Tasks = methods
	return
}

func (u *UserBehaviour) Execute() {
	fooVal := reflect.ValueOf(u.Behaviour)
	for _, m := range u.Tasks {
		res := fooVal.MethodByName(m).Call([]reflect.Value{})
		response := res[0].Interface().(*http.Response)
		url := response.Request.URL
		statusCode := response.StatusCode
		fmt.Println(url, statusCode)
	}
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
	Wasp = &UserBehaviour{
		custUser,
		nil,
	}
	fmt.Println(Wasp.GetMethods())
	Wasp.Execute()
}
