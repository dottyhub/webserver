package webserver

import (
	"net/http"
)

type Multiplexer interface {
	Register(path string, view View, controller Controller)
	GetHandler() http.Handler
}
