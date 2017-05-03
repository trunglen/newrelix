package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
"newrelixdemo/wapii"
	newrelic "github.com/newrelic/go-agent"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func main() {
	var server = http.NewServeMux()
	var apiServer = apii.NewApiServer()
	server.Handle("/api/", http.StripPrefix("/api", apiServer))
	// Create a config.  You need to provide the desired application name
	// and your New Relic license key.
	cfg := newrelic.NewConfig("Example App", "101549ce7776071f41c36184c4e7ded4df588948")

	// Create an application.  This represents an application in the New
	// Relic UI.
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Wrap helloHandler.  The performance of this handler will be recorded.
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", helloHandler))
	http.ListenAndServe(":8001", nil)
}
