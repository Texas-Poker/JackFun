package server_lobby

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/logger"
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
	//如果从redis里获取的用户信息为零值，说明这个用户没有创建过，此时创建一个新的用户数据结构
	if u.Id.Get()==0{
		u.Id.Set(uid)
		u.NickName.Set(fmt.Sprintf("玩家%s",token))
		u.Sex.Set(int64(pb_enum.Sex_Unknow))
		u.Diamond.Set(1000)
		u.Gold.Set(2000)
		u.Lv.Set(1)
	}
	u.ManualSave()
	logger.Log.Infof("new user to session, user=%+v\n", u)
	pojo.SetUserToSession(s, u)

	return &pb_lobby.RespAuth{ErrCode: pb_enum.ErrorCode_OK}, nil
}

