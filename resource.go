package webserver

import "net/http"

type View interface {
	Get() http.HandlerFunc
}

type Controller interface {
	Post() http.HandlerFunc
	Put() http.HandlerFunc
	Delete() http.HandlerFunc
}

type Resource interface {
	View() View
	Controller() Controller
}
