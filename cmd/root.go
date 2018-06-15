package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DudeWhoCode/yellowjacket/backend"
	"github.com/DudeWhoCode/yellowjacket/server"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "yellowjacket",
	Short: "YellowJacket is a simple, distributed load testing framework",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		startApp()
	},
}

func Execute() {
	RootCmd.PersistentFlags().StringVarP(&backend.WebHost, "web-host", "w", "",
		"Address of the webapp")
	RootCmd.PersistentFlags().StringVarP(&backend.PortNo, "port", "p", "",
		"Port of the webapp")
	RootCmd.PersistentFlags().StringVarP(&backend.TargetHost, "host", "H", "",
		"Host to target")
	fmt.Println("All flags set")
	if err := RootCmd.Execute(); err != nil {
		log.Fatal("Command Execution error : ", err)
		os.Exit(-1)
	}
}

func startApp() {
	rt := server.NewRouter()
	handler := cors.Default().Handler(rt)
	// Setup the file server to serve the static site with css and js
	fs := http.FileServer(http.Dir("templates/"))
	rt.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", fs))
	// Start the SSE broker
	b := backend.GetBroker()
	b.Start()
	webAddr := fmt.Sprintf("%s:%s", backend.WebHost, backend.PortNo)
	log.Printf("YellowJacket is running in %s", webAddr)
	log.Fatal(http.ListenAndServe(webAddr, handler))
}
