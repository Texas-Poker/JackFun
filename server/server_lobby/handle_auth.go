package server_lobby

import (
	"context"
	"github.com/topfreegames/pitaya"
	"server/pb/pb_enum"
	"server/pb/pb_lobby"
	"strconv"
)

func (this *ComponentLobby) ReqAuth(ctx context.Context, req *pb_lobby.ReqAuth) (*pb_lobby.RespAuth, error) {
	token:=req.GetToken()
	isAlreadyLogin,uid,err:= this.db_login_handler.IsAlreadyLogin(token)
	if err!=nil || !isAlreadyLogin{
		return &pb_lobby.RespAuth{ErrCode: pb_enum.ErrorCode_AuthFailed}, nil
	}
	s := pitaya.GetSessionFromCtx(ctx)
	s.Bind(ctx,strconv.FormatInt(uid,10))

	return &pb_lobby.RespAuth{ErrCode: pb_enum.ErrorCode_OK}, nil
}