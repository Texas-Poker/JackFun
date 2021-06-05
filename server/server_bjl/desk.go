package server_bjl

import (
	"github.com/looplab/fsm"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/timer"
	"server/pb/pb_bjl"
	"server/server_bjl/poker"
	"time"
)

type Desk struct {
	deskID        string            //桌子的ID
	deskLevel     pb_bjl.BjlLevel   //桌子的游戏场次级别
	deskGroupKey  string            //桌子group的key
	deskRound     uint32            //桌子的第几轮
	allPoker      *poker.PokerBjl   //所有的扑克牌
	deskFSM       *fsm.FSM          //桌子的状态机
	timer         *timer.Timer      //状态机的计时器
	minBet        uint32            //最小下注额度
	maxBet        uint32            //最大下注额度
	roundBetInfo  []*pb_bjl.BetInfo //桌子上当前局的下注流水
	historyResult []*pb_bjl.Result  //桌子上历史结果
	curDeskPoker *deskPoker			//当前桌子上发的牌
}


func NewDesk(level pb_bjl.BjlLevel, deskID string) *Desk {
	desk := &Desk{
		deskLevel:     level,
		deskID:        deskID,
		roundBetInfo:  make([]*pb_bjl.BetInfo, 0),
		historyResult: make([]*pb_bjl.Result, 0),
	}
	desk.allPoker = poker.NewPokerBjl(52)

	switch level {
	case pb_bjl.BjlLevel_Primary:
		desk.minBet = 2
		desk.maxBet = 5
	case pb_bjl.BjlLevel_Middle:
		desk.minBet = 2
		desk.maxBet = 5
	case pb_bjl.BjlLevel_High:
		desk.minBet = 2
		desk.maxBet = 5
	case pb_bjl.BjlLevel_King:
		desk.minBet = 2
		desk.maxBet = 5
	}
	return desk
}

//Start 启动桌子的状态机 startTime为各桌子错峰启动时间，错峰启动防止同一时间多个桌子一起运算
func (this *Desk) Start(startTime int64) {
	//初始化状态机
	this.deskFSM = fsm.NewFSM(
		stateNone,
		fsm.Events{
			{Name: stateReady, Src: []string{stateNone, stateSettle}, Dst: stateReady},
			{Name: stateBet, Src: []string{stateReady}, Dst: stateBet},
			{Name: stateSend, Src: []string{stateBet}, Dst: stateSend},
			{Name: stateShow, Src: []string{stateSend}, Dst: stateShow},
			{Name: stateSettle, Src: []string{stateShow}, Dst: stateSettle},
		},
		fsm.Callbacks{
			enterStateAny: func(e *fsm.Event) { this.OnEnterStatusAny(e) },

			enterStateReady:  func(e *fsm.Event) { this.OnEnterStatusReady(e) },
			enterStateBet:    func(e *fsm.Event) { this.OnEnterStatusBet(e) },
			enterStateSend:   func(e *fsm.Event) { this.OnEnterStatusSend(e) },
			enterStateShow:   func(e *fsm.Event) { this.OnEnterStatusShow(e) },
			enterStateSettle: func(e *fsm.Event) { this.OnEnterStatusSettle(e) },
		},
	)

	//时间标识位，单位（毫秒）
	timeTag := time.Now().UnixNano()/1e6
	this.timer = pitaya.NewTimer(time.Second, func() {
		nowTag := time.Now().UnixNano()/1e6
		switch this.deskFSM.Current() {
		case stateNone:
			if nowTag-timeTag > startTime*1000 {
				timeTag = nowTag
				this.deskFSM.Event(stateReady)
			}
		case stateReady:
			if nowTag-timeTag > 3*1000 {
				timeTag = nowTag
				this.deskFSM.Event(stateBet)
			}
		case stateBet:
			if nowTag-timeTag > 10*1000 {
				timeTag = nowTag
				this.deskFSM.Event(stateSend)
			}
		case stateSend:
			if nowTag-timeTag > 5*1000 {
				timeTag = nowTag
				this.deskFSM.Event(stateShow)
			}
		case stateShow:
			if nowTag-timeTag > 5*1000 {
				timeTag = nowTag
				this.deskFSM.Event(stateSettle)
			}
		case stateSettle:
			if nowTag-timeTag > 5*1000 {
				this.deskFSM.Event(stateReady)
			}
		}

	})
}

func (this *Desk) OnPlayerBet(uid int64, betInfo *pb_bjl.NotifyBet) {
	this.roundBetInfo = append(this.roundBetInfo, &pb_bjl.BetInfo{
		UID:   uid,
		Area:  betInfo.Area,
		Count: betInfo.Count,
	})
}

