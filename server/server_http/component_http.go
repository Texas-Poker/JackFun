package server_http

import (
	"github.com/topfreegames/pitaya/component"
	"log"
	"net/http"
)

type ComponentHttp struct {
	component.Base
}

func (this *ComponentHttp) Init() {
	this.startHttpServer()
}

func NewComponentHttp() *ComponentHttp {
	return &ComponentHttp{}
}

func (this *ComponentHttp) startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8088"}
	http.HandleFunc("/entry", this.entry)
	http.HandleFunc("/login", this.login)
	http.HandleFunc("/test", this.test)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	return srv
}


