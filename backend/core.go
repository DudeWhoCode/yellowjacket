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

func (s *swarm) WebInputs(waspCnt, hatchRte int) *swarm {
	s.WaspCount = waspCnt
	s.HatchRate = hatchRte
	return s
}

func (s *swarm) FileInputs(input *UserBehaviour) *swarm {
	s.Wasp = input
	return s
}

func (s *swarm) SetChan(outputChan chan RawResponse) *swarm {
	s.OutChan = outputChan
	return s
}

func (s *swarm) UpdateStatus(numReq, numFail int) *swarm {
	s.NumReq = numReq
	s.NumFail = numFail
	return s
}

var sw *swarm
var once sync.Once

func GetSwarm() *swarm {
	once.Do(func() {
		sw = &swarm{}
	})
	return sw
}

func (u *UserBehaviour) GetMethods() (methods []string) {
	fooType := reflect.TypeOf(u.Behaviour)
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		methods = append(methods, method.Name)
	}
	u.Tasks = methods
	return
}

func (u *UserBehaviour) ExecuteOne(methodName string) (response *http.Response) {
	fooVal := reflect.ValueOf(u.Behaviour)
	res := fooVal.MethodByName(methodName).Call([]reflect.Value{})
	response = res[0].Interface().(*http.Response)
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
	wasp := &UserBehaviour{
		custUser,
		nil,
	}
	s := GetSwarm()
	s.FileInputs(wasp)
}
