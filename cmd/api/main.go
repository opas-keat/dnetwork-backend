package main

import (
	"ect/dnetwork/backend/pkg/app"
	"ect/dnetwork/backend/pkg/config"
	"ect/dnetwork/backend/pkg/module"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
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
	// Load acl model and policy
	enforcer, err := casbin.NewEnforcer("config/acl_model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	app := app.New(cfg)
	// Cleanup when server stopped
	defer app.Close()

	// For Liveness Probe
	app.CreateLivenessFile()

	// Initialize data sources
	app.InitDS()

	// Create router (mux/gin/fiber)
	app.InitRouter(enforcer)

	// Initialize module with dependency injection
	module.Init(app.Context)
	// Start server
	app.ServeHTTP()
}

func GreetingAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher")
}
