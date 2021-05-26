package server_http

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/bwmarrin/snowflake"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/logger"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"server/dao/pojo"
	"server/pb/pb_enum"
	"server/pb/pb_http"
	"server/pb/pb_lobby"
	"strings"
)

//客户端与服务端连接的密钥
var key = "天王盖地虎"

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//entry 客户端与服务器连接的第一个方法，这个方法用来获取
func (this *ComponentHttp) entry(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	req := new(pb_http.ReqEntry)
	if err := proto.Unmarshal(buf, req); err != nil {
		return
	}
	log.Printf("[entry], req.Secret=%s\n", req.Secret)
	resp := new(pb_http.RespEntry)
	//如果客户端的包里不带密钥或是密钥错误，将无法获取真实的游戏服务器地址
	if strings.Contains(req.Secret, key) && strings.Contains(req.Secret, "宝塔镇河妖") {
		resp.ErrCode = pb_enum.ErrorCode_OK
		resp.LoginUrl = "http://127.0.0.1:8088/login"
		resp.RegisterUrl = "http://127.0.0.1:8088/register"
		resp.TcpUrl = "127.0.0.1:3250"
	} else {
		resp.ErrCode = pb_enum.ErrorCode_EntryError
	}

	if pbByte, err := proto.Marshal(resp); err == nil {
		log.Printf("[entry] result=%v\n", resp)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(pbByte)
	}
}

func (this *ComponentHttp) login(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	req := new(pb_http.ReqLogin)
	if err := proto.Unmarshal(buf, req); err != nil {
		return
	}

	log.Printf("onLogin account=%s, password=%s\n", req.Account, req.Password)

	resp := new(pb_http.RespLogin)
	//判断是否是空找得到数据
	if isExist, _ := pojo.CheckAccountExist(req.Account); !isExist {
		resp.ErrCode = pb_enum.ErrorCode_Default
	} else {
		if password, _, err := pojo.QueryAccount(req.Account); err == nil {
			if req.Password == password {
				//将account+时间生成token
				resp.ErrCode = pb_enum.ErrorCode_OK
			} else {
				resp.ErrCode = pb_enum.ErrorCode_LoginAccountOrPasswordError
			}
		}
	}

	logger.Log.Infof("[entry] result=%v\n", resp)
	if pbByte, err := proto.Marshal(resp); err == nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(pbByte)
	}
}

func (this *ComponentHttp) register(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadAll(r.Body)
	req := new(pb_http.ReqRegister)
	if err := proto.Unmarshal(buf, req); err != nil {
		return
	}

	log.Printf("onRegister account=%s, password=%s\n", req.Account, req.Password)

	resp := new(pb_http.RespRegister)
	//判断是否是空找得到数据
	if isExist, _ := pojo.CheckAccountExist(req.Account); isExist {
		resp.ErrCode = pb_enum.ErrorCode_RegisterAccountExit
	} else {
		n, err := snowflake.NewNode(1)
		if err != nil {
			println(err)
		}
		id := n.Generate().Int64()
		pojo.NewUserAndSaveRedis(req.Account, req.Password, id)
		resp.ErrCode = pb_enum.ErrorCode_OK
	}

	logger.Log.Infof("[register] result=%v\n", resp)
	if pbByte, err := proto.Marshal(resp); err == nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(pbByte)
	}
}

func (this *ComponentHttp) test(w http.ResponseWriter, r *http.Request) {
	req := &pb_lobby.ReqLobbyInfo{

	}
	resp := new(pb_lobby.RespLobbyInfo)
	if err := pitaya.RPC(context.Background(), "ServerLobby.ComponentLobby.Test", resp, req); err != nil {

		return
	}
	log.Printf("rpc result, resp.errorcode=%s\n", resp.ErrCode)
	respByte, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	w.Write(respByte)
}
