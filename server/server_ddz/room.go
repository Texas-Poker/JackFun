package server_ddz

import (
	"context"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/config"
	"github.com/topfreegames/pitaya/groups"
	"github.com/topfreegames/pitaya/timer"
	"server/dao/pojo"
	"server/pb/pb_ddz"
	"time"
)

type (
	Room struct {
		component.Base
		roomGroupKey string //房间group的key
		players      map[int]*Player
		timer        *timer.Timer
		Stats        *Stats
		desks []*Desk
	}

	Stats struct {
		outboundBytes int
		inboundBytes  int
	}
)

func (this *Room) Init() {
	gsi := groups.NewMemoryGroupService(config.NewConfig())
	pitaya.InitGroups(gsi)
	this.roomGroupKey = "group_room_ddz"
	pitaya.GroupCreate(context.Background(), this.roomGroupKey)
	this.desks = make([]*Desk, 0)
	for i := 0; i < 10; i++ {
		desk := NewDesk(1000+i)
		this.desks = append(this.desks, desk)
	}
}

func (this *Room) AfterInit() {
	this.timer = pitaya.NewTimer(time.Minute, func() {
		count, err := pitaya.GroupCountMembers(context.Background(), this.roomGroupKey)
		println("UserCount: Time=>", time.Now().String(), "Count=>", count, "Error=>", err)
		println("OutboundBytes", this.Stats.outboundBytes)
		println("InboundBytes", this.Stats.outboundBytes)
	})
}

// JoinDDZ 加入斗地主房间
func (this *Room) JoinDDZ(ctx context.Context, req *pb_ddz.ReqJoinRoom) (*pb_ddz.RespJoinRoom, error) {
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)
	s := pitaya.GetSessionFromCtx(ctx)
	err := pitaya.GroupAddMember(ctx, this.roomGroupKey, s.UID())
	if err != nil {
		logger.Error("Failed to join room")
		logger.Error(err)
		return nil, err
	}

	u := pojo.GetUserFromSession(s)
	broadcastInfo := &pb_ddz.SyncPlayerJoinRoom{
		PlayerId: u.UID.Get(),
		NickName: u.Char.NickName.Get(),
	}
	err = pitaya.GroupBroadcast(ctx, "connector", this.roomGroupKey, "SeverDdz.ComponentRoom.SyncPlayerJoinRoom", broadcastInfo)
	if err != nil {
		logger.Error("Failed to broadcast onNewUser")
		logger.Error(err)
		return nil, err
	}
	resp := &pb_ddz.RespJoinRoom{
		Info: []*pb_ddz.RespJoinRoom_DeskInfo{
			{DeskId: 1001, Count: 2},
			{DeskId: 1002, Count: 0},
			{DeskId: 1003, Count: 1},
			{DeskId: 1004, Count: 3},
			{DeskId: 1005, Count: 0},
		},
	}
	return resp, nil
}
