package server_http

import (
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"log"
	"net/http"
	"server/dao/redis_module"
)

type ComponentHttp struct {
	component.Base
	redisModule *redis_module.RedisModule

}

func (this *ComponentHttp) Init() {
	this.startHttpServer()
	if module, err := pitaya.GetModule("redisModule"); err == nil {
		this.redisModule = module.(*redis_module.RedisModule)
	}


}

func NewComponentHttp() *ComponentHttp {
	return &ComponentHttp{}
}

func (this *ComponentHttp) startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8088"}
	http.HandleFunc("/entry", this.entry)
	http.HandleFunc("/register", this.register)
	http.HandleFunc("/login", this.login)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	return srv
}
