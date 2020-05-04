package webserver

type PlugIn interface {
	Init(root string, mux Multiplexer)
}
