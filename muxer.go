package webserver

import (
	"net/http"
)

type Multiplexer interface {
	RegisterFunc(path string, f func(http.ResponseWriter, *http.Request), method string)
}
