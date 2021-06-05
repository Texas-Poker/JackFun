package server_ddz

import (
	"fmt"
	"github.com/looplab/fsm"
	_ "github.com/looplab/fsm"
	"github.com/topfreegames/pitaya/groups"
	"github.com/topfreegames/pitaya/session"
	"server/server_ddz/poker"
)

type Desk struct {
	deskID         int
	deskGroupKey   string
	players        *Player     //所有的玩家
	allPoker       poker.Poker //所有的扑克牌
	group          groups.GroupService
	dzSeat         int   //地主的坐位号
	nextPokerIndex int   //下一张poker在整副牌数组中的索引
	lastHitPokers  []int //最后打出的牌的集合
	lastHitPlayer  int   //最后一个出牌的玩家ID
	lastTipPlayer  int   //最后一次提示出牌的玩家的ID

	fsm *fsm.FSM
}

func NewDesk(id int) *Desk {
	desk := &Desk{
		deskID:       id,
		deskGroupKey: fmt.Sprintf("group_desk_%d", id),
	}
	desk.fsm = fsm.NewFSM(
		"status_ready",
		fsm.Events{
			{Name: "status_", Src: []string{"status_ready"}, Dst: "join"},
		},
		fsm.Callbacks{},
	)

	return desk
}

// 玩家数量
func (d *Desk) totalPlayerCount() int {
	return 3
}

// 如果是重新进入 isReJoin: true
func (d *Desk) JoinDesk(s *session.Session) error {
	//uid := s.UID()
	//var (
	//	p   *Player
	//	err error
	//)
	//
	//if isReJoin {
	//	d.dissolve.updateOnlineStatus(uid, true)
	//	p, err = d.playerWithId(uid)
	//	if err != nil {
	//		d.logger.Errorf("玩家: %d重新加入房间, 但是没有找到玩家在房间中的数据", uid)
	//		return err
	//	}
	//
	//	// 加入分组
	//	d.group.Add(s)
	//} else {
	//	exists := false
	//	for _, p := range d.players {
	//		if p.Uid() == uid {
	//			exists = true
	//			p.logger.Warn("玩家已经在房间中")
	//			break
	//		}
	//	}
	//	if !exists {
	//		p = s.Value(kCurPlayer).(*Player)
	//		d.players = append(d.players, p)
	//		for i, p := range d.players {
	//			p.setDesk(d, i)
	//		}
	//		d.roundStats[uid] = &history.Record{}
	//	}
	//}

	return nil
}
