package main

import (
	"fmt"
	"net/http"

	"github.com/dottyhub/webserver/gorilla"
)

func main() {
	mupl := gorilla.NewMultiplexer()
	mupl.RegisterFunc("/", homeHandler, http.MethodGet)

	http.ListenAndServe(":3000", mupl.GetHandler())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}
