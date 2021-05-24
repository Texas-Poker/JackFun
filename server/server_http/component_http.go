package server_http

import (
	"context"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"log"
	"net/http"
	"server/pb/pb_lobby"
)

type ComponentHttp struct {
	component.Base
}

func (this *ComponentHttp) Init() {
	log.Println("sv http init")
	this.startHttpServer()
}

func NewComponentHttp() *ComponentHttp {
	return &ComponentHttp{}
}

func (this *ComponentHttp) startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8088"}
	http.HandleFunc("/entry", this.entry)
	http.HandleFunc("/login", this.login)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	return srv
}

func (this *ComponentHttp) entry(w http.ResponseWriter, r *http.Request) {

}

func (this *ComponentHttp) login(w http.ResponseWriter, r *http.Request) {
	req := &pb_lobby.ReqLogin{
		Account:  "abc",
		Password: "123",
	}
	resp := &pb_lobby.RespLogin{}
	if err := pitaya.RPC(context.Background(), "sv_lobby.component_lobby.ReqLogin", resp, req); err != nil {
		log.Println("rpc call sv_lobby.cp_login.reqlogin err, err=",err)
		return
	}
	log.Printf("---------------error code=%s\n--------------", resp.ErrCode.String())
}
