package main

import (
	"fmt"
	"net/http"

	"github.com/dottyhub/webserver/gorilla"
)

type myController struct {
}

func (f *myController) Post() http.HandlerFunc {
	return nil
}

func (f *myController) Put() http.HandlerFunc {
	return nil
}

func (f *myController) Delete() http.HandlerFunc {
	return nil
}

type myView struct {
}

func (f *myView) Get() http.HandlerFunc {
	return f.homeHandler
}

func (f *myView) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func main() {
	mupl := gorilla.NewMultiplexer()
	mupl.Register("/", &myView{}, &myController{})

	http.ListenAndServe(":3000", mupl.GetHandler())
}
