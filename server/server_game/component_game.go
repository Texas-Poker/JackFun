package server_game

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"server/pb/pb_game"
)



type componentGame struct {
	component.Base
}

type SessionData struct {
	Data map[string]interface{}
}

func (c *componentGame) GetSessionData(ctx context.Context) (*SessionData, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	res := &SessionData{
		Data: s.GetData(),
	}
	return res, nil
}

func (c *componentGame) SetSessionData(ctx context.Context, data *SessionData) (*pb_game.SyncGateLinkState, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	err := s.SetData(data.Data)
	if err != nil {
		return nil, pitaya.Error(err, "CN-000", map[string]string{"failed": "set data"})
	}
	return &pb_game.SyncGateLinkState{ErrorCode: "success"}, nil
}

func (c *componentGame) NotifySessionData(ctx context.Context, data *SessionData) {
	s := pitaya.GetSessionFromCtx(ctx)
	err := s.SetData(data.Data)
	if err != nil {
		fmt.Println("got error on notify", err)
	}
}

func (c *componentGame) RemoteFunc(ctx context.Context, message []byte) (*pb_game.SyncGateLinkState, error) {
	fmt.Printf("received a remote call with this message: %s\n", message)
	return &pb_game.SyncGateLinkState{ErrorCode: "success"}, nil
}
