package server_bjl

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/config"
	"github.com/topfreegames/pitaya/groups"
	"server/dao/pojo"
	"server/pb/pb_bjl"
	"server/pb/pb_enum"
)

type (
	Bjl struct {
		component.Base
		groupKey string //游戏group的key
		desks    []*Desk
	}
)

func NewComponentBjl() *Bjl {
	return &Bjl{

	}
}

func (this *Bjl) Init() {
	gsi := groups.NewMemoryGroupService(config.NewConfig())
	pitaya.InitGroups(gsi)
	this.groupKey = "group_room_bjl"
	pitaya.GroupCreate(context.Background(), this.groupKey)

	//初始化4*4共计16张桌子
	this.desks = []*Desk{
		NewDesk(pb_bjl.BjlLevel_Primary, "2000_1_1"),
		//NewDesk(pb_bjl.BjlLevel_Primary, "2000_1_2"),
		//NewDesk(pb_bjl.BjlLevel_Primary, "2000_1_3"),
		//NewDesk(pb_bjl.BjlLevel_Primary, "2000_1_4"),

		//NewDesk(pb_bjl.BjlLevel_Middle, "2000_2_1"),
		//NewDesk(pb_bjl.BjlLevel_Middle, "2000_2_2"),
		//NewDesk(pb_bjl.BjlLevel_Middle, "2000_2_3"),
		//NewDesk(pb_bjl.BjlLevel_Middle, "2000_2_4"),

		//NewDesk(pb_bjl.BjlLevel_High, "2000_3_1"),
		//NewDesk(pb_bjl.BjlLevel_High, "2000_3_2"),
		//NewDesk(pb_bjl.BjlLevel_High, "2000_3_3"),
		//NewDesk(pb_bjl.BjlLevel_High, "2000_3_4"),
		//
		//NewDesk(pb_bjl.BjlLevel_King, "2000_4_1"),
		//NewDesk(pb_bjl.BjlLevel_King, "2000_4_2"),
		//NewDesk(pb_bjl.BjlLevel_King, "2000_4_3"),
		//NewDesk(pb_bjl.BjlLevel_King, "2000_4_4"),
	}

}

func (this *Bjl) AfterInit() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//randomTime := r.Int63n(5)
	for i, desk := range this.desks {
		desk.Start(int64(i))
	}
}

func (this *Bjl) CallJoinRoom(ctx context.Context, req *pb_bjl.ReqJoinRoom) (*pb_bjl.RespJoinRoom, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	u := pojo.GetUserFromSession(s)
	if err := pitaya.GroupAddMember(ctx, this.groupKey, s.UID()); err != nil {
		return nil, err
	}
	u.Lobby.AtRoomID.Set(1004)
	resp := new(pb_bjl.RespJoinRoom)
	for _, desk := range this.desks {
		playerCount, _ := pitaya.GroupCountMembers(ctx, desk.deskGroupKey)
		resp.Info = append(resp.Info, &pb_bjl.RespJoinRoom_DeskInfo{
			Level:       desk.deskLevel,
			DeskId:      desk.deskID,
			PlayerCount: uint32(playerCount),
		})
	}
	return resp, nil
}

func (this *Bjl) CallJoinDesk(ctx context.Context, req *pb_bjl.ReqJoinDesk) (*pb_bjl.RespJoinDesk, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	u := pojo.GetUserFromSession(s)
	var targetDesk *Desk = nil
	for _, desk := range this.desks {
		if desk.deskID == req.DeskId {
			targetDesk = desk
		}
	}
	if targetDesk == nil {
		return nil, pitaya.Error(errors.New(pb_enum.ErrorCode_DeskUnExistent.String()), uuid.New().String())
	}
	playerCount, _ := pitaya.GroupCountMembers(ctx, targetDesk.deskGroupKey)
	if playerCount >= 100 {
		return nil, pitaya.Error(errors.New(pb_enum.ErrorCode_DeskPlayerFull.String()), uuid.New().String())
	}

	u.Bjl.AtDeskID.Set(targetDesk.deskID)
	resp := new(pb_bjl.RespJoinDesk)
	resp.GameStatus = targetDesk.deskFSM.Current()
	resp.BetInfos = targetDesk.curRoundBetInfo
	return resp, nil
}

func (this *Bjl) NotifyBet(ctx context.Context, req *pb_bjl.NotifyBet) error  {
	s := pitaya.GetSessionFromCtx(ctx)
	u := pojo.GetUserFromSession(s)
	targetDeskId:=u.Bjl.AtDeskID.Get()
	var targetDesk *Desk = nil
	for _, desk := range this.desks {
		if desk.deskID == targetDeskId {
			targetDesk = desk
		}
	}
	if targetDesk == nil {
		return  pitaya.Error(errors.New(pb_enum.ErrorCode_DeskUnExistent.String()), uuid.New().String())
	}

	targetDesk.OnPlayerBet(u.UID.Get(), req)

	return nil
}