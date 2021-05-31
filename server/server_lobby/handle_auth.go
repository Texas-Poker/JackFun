package server_lobby

import (
	"context"
	"github.com/topfreegames/pitaya"
	"server/dao/pojo"
	"server/pb/pb_enum"
	"server/pb/pb_lobby"
	"strconv"
)

func (this *ComponentLobby) ReqAuth(ctx context.Context, req *pb_lobby.ReqAuth) (*pb_lobby.RespAuth, error) {
	token := req.GetToken()
	isAlreadyLogin, uid, err := this.redisModule.IsAlreadyLogin(token)
	if err != nil || !isAlreadyLogin {
		return &pb_lobby.RespAuth{ErrCode: pb_enum.ErrorCode_AuthFailed}, nil
	}
	s := pitaya.GetSessionFromCtx(ctx)

	s.Bind(ctx, strconv.FormatInt(uid, 10))

	dbUser, err := this.redisModule.GetUserByUId(uid)
	if err != nil {
		panic(err)
	}

	u := new(pojo.User)
	err = u.Load(u, dbUser)
	if err != nil {
		panic(err)
	}
	pojo.SetUserToSession(s, u)

	return &pb_lobby.RespAuth{
		ErrCode:  pb_enum.ErrorCode_OK,
		UID:      u.UID.Get(),
		NickName: u.Char.NickName.Get(),
		Sex:      pb_enum.Sex(u.Char.Sex.Get()),
		Lv:       u.Char.Lv.Get(),
		Gold:     u.Char.Gold.Get(),
		Diamond:  u.Char.Diamond.Get(),
	}, nil
}
