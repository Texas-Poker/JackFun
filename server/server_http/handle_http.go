package server_http

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/topfreegames/pitaya"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"server/pb/pb_common"
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
	buf, err := ioutil.ReadAll(r.Body)
	req := new(pb_http.ReqEntry)
	if err := proto.Unmarshal(buf, req); err != nil {
		return
	}
	log.Printf("[entry], req.Secret=%s\n", req.Secret)
	resp := new(pb_http.RespEntry)
	//如果客户端的包里不带密钥或是密钥错误，将无法获取真实的游戏服务器地址
	if strings.Contains(req.Secret, key) && strings.Contains(req.Secret, "宝塔镇河妖") {
		resp.ErrCode = pb_common.ErrorCode_OK
		resp.LoginUrl = "http://127.0.0.1:8088/Login"
		resp.RegisterUrl = "http://127.0.0.1:8088/Register"
		resp.TcpUrl = "127.0.0.1:3250"
	} else {
		resp.ErrCode = pb_common.ErrorCode_EntryError
	}

	pbByte, err := proto.Marshal(resp)
	log.Printf("[entry] result=%v\n", resp)
	if err != nil {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(pbByte)
}

func (this *ComponentHttp) login(w http.ResponseWriter, r *http.Request) {

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
	//json.Marshal()
	if err != nil {
		return
	}
	w.Write(respByte)
}
