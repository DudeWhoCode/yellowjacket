package main

import (
	"github.com/DudeWhoCode/yellowjacket/backend"
	"github.com/DudeWhoCode/yellowjacket/cmd"
)

func main() {
	backend.LoadModule()
	cmd.Execute()
}
