package gorilla

import (
	"net/http"

	"github.com/dottyhub/webserver"
	"github.com/gorilla/mux"
)

type multiplexer struct {
	instance *mux.Router
}

func NewMultiplexer() webserver.Multiplexer {
	return &multiplexer{instance: mux.NewRouter()}
}

func (mupl *multiplexer) RegisterFunc(path string, f func(http.ResponseWriter, *http.Request), method string) {
	mupl.instance.HandleFunc(path, f).Methods(method)
}

func (mupl *multiplexer) GetHandler() http.Handler {
	return mupl.instance
}
