package main

import (
	"ect/dnetwork/backend/pkg/app"
	"ect/dnetwork/backend/pkg/config"
	"ect/dnetwork/backend/pkg/module"
	"fmt"
	"net/http"
)

var version = "dev"

func Greeting() string {
	return "Hello"
}

func GoodBye(name string) string {
	text := "Goodbye, " + name + "..."
	return text
}

func Add(a int, b int) int {
	return a + b
}

func main() {
	// http.HandleFunc("/", GreetingAPI)
	// http.ListenAndServe(":3000", nil)

	// Load config
	cfg := config.LoadConfig()

	app := app.New(cfg)
	// Cleanup when server stopped
	defer app.Close()

	// For Liveness Probe
	app.CreateLivenessFile()

	// Initialize data sources
	app.InitDS()

	// Create router (mux/gin/fiber)
	app.InitRouter()

	// Initialize module with dependency injection
	module.Init(app.Context)

	// Start server
	app.ServeHTTP()
}

func GreetingAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher")
}
