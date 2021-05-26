package server_lobby

import (
	"context"
	"server/pb/pb_enum"
	"server/pb/pb_lobby"
)

func (this *ComponentLobby) ReqLogin(ctx context.Context, req *pb_lobby.ReqLobbyInfo) (*pb_lobby.RespLobbyInfo, error) {
	return &pb_lobby.RespLobbyInfo{ErrCode: pb_enum.ErrorCode_AuthFailed}, nil
}