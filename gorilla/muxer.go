package gorilla

import (
	"net/http"

	"github.com/dottyhub/webserver"
	"github.com/gorilla/mux"
)

type multiplexer struct {
	instance *mux.Router
}

func NewMultiplexer() (webserver.Multiplexer, error) {
	return &multiplexer{instance: mux.NewRouter()}, nil
}

func (mupl *multiplexer) RegisterFunc(path string, f func(http.ResponseWriter, *http.Request), method string) {
	mupl.instance.HandleFunc(path, f).Methods(method)
}
