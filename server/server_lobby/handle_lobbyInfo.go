package server_lobby

import (
	"context"
	"server/pb/pb_lobby"
)

func (this *ComponentLobby) CallLobbyInfo(ctx context.Context, req *pb_lobby.ReqLobbyInfo) (*pb_lobby.RespLobbyInfo, error) {
	resp := &pb_lobby.RespLobbyInfo{
		Infos: []*pb_lobby.RespLobbyInfo_LobbyInfo{
			{GameId: 1001, GameName: "斗地主", IsOpen: true},
			{GameId: 1002, GameName: "德州扑克", IsOpen: false},
			{GameId: 1003, GameName: "炸金花", IsOpen: false},
			{GameId: 1004, GameName: "百家乐", IsOpen: true},
		}}
	return resp, nil
}
