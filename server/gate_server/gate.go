package gate_server

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"server/pb/pb_gate"
)

type GateRemote struct {
	component.Base
}

type Gate struct {
	component.Base
}


type SessionData struct {
	Data map[string]interface{}
}


func (c *Gate) GetSessionData(ctx context.Context) (*SessionData, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	res := &SessionData{
		Data: s.GetData(),
	}
	return res, nil
}


func (c *Gate) SetSessionData(ctx context.Context, data *SessionData) (*pb_gate.SyncGateLinkState, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	err := s.SetData(data.Data)
	if err != nil {
		return nil, pitaya.Error(err, "CN-000", map[string]string{"failed": "set data"})
	}
	return &pb_gate.SyncGateLinkState{ErrorCode: "success"}, nil
}


func (c *Gate) NotifySessionData(ctx context.Context, data *SessionData) {
	s := pitaya.GetSessionFromCtx(ctx)
	err := s.SetData(data.Data)
	if err != nil {
		fmt.Println("got error on notify", err)
	}
}


func (c *Gate) RemoteFunc(ctx context.Context, message []byte) (*pb_gate.SyncGateLinkState, error) {
	fmt.Printf("received a remote call with this message: %s\n", message)
	return &pb_gate.SyncGateLinkState{ErrorCode: "success"}, nil
}
