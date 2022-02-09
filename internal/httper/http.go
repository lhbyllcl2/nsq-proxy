package httper

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/rakyll/statik/fs"
	_ "nsq-proxy/internal/statik"
)

type Httper struct {
	addr     string
	server   *http.Server
	statikFS http.FileSystem
}

func NewHttper(addr string) *Httper {
	statikFS, err := fs.New()
	if err != nil {
		panic("NewHttper statikFS error: " + err.Error())
	}
	return &Httper{
		addr:     addr,
		server:   &http.Server{Addr: addr, Handler: nil},
		statikFS: statikFS,
	}
}

// Run 启动HTTP
func (h *Httper) Run() {
	h.router()
	go func() {
		err := h.server.ListenAndServe()
		if err != nil {
			panic("ListenAndServe error: " + err.Error())
		}
	}()
}
