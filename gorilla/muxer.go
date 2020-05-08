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

func (mupl *multiplexer) Register(path string, view webserver.View, controller webserver.Controller) {
	if view != nil && view.Get() != nil {
		mupl.instance.HandleFunc(path, view.Get()).Methods(http.MethodGet)
	}
	if controller != nil && controller.Post() != nil {
		mupl.instance.HandleFunc(path, controller.Post()).Methods(http.MethodPost)
	}
	if controller != nil && controller.Put() != nil {
		mupl.instance.HandleFunc(path, controller.Put()).Methods(http.MethodPut)
	}
	if controller != nil && controller.Delete() != nil {
		mupl.instance.HandleFunc(path, controller.Delete()).Methods(http.MethodDelete)
	}
}

func (mupl *multiplexer) GetHandler() http.Handler {
	return mupl.instance
}
