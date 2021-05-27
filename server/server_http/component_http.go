package server_http

import (
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"log"
	"net/http"
	"server/dao/db_handler"
)

type ComponentHttp struct {
	component.Base
	db_reigster_handler *db_handler.RedisRegisterHandler
	db_login_handler  *db_handler.RedisLoginHandler
}

func (this *ComponentHttp) Init() {
	this.startHttpServer()
	if dbHandler, err := pitaya.GetModule("dbHandlerRegister"); err == nil {
		this.db_reigster_handler = dbHandler.(*db_handler.RedisRegisterHandler)
	}

	if dbHandler, err := pitaya.GetModule("dbLoginRegister"); err == nil {
		this.db_login_handler = dbHandler.(*db_handler.RedisLoginHandler)
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
