package server_bjl

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/logger"
	"server/pb/pb_bjl"
)

const (
	stateNone   = "stateNone"   //空阶段
	stateReady  = "stateReady"  //准备阶段
	stateBet    = "stateBet"    //下注阶段
	stateSend   = "stateSend"   //发牌阶段
	stateShow   = "stateShow"   //显示牌阶段
	stateSettle = "stateSettle" //结算阶段
)

const (
	enterStateAny = "enter_state" //进入任意阶段

	enterStateReady  = "enter_stateReady"
	enterStateBet    = "enter_stateBet"
	enterStateSend   = "enter_stateSend"
	enterStateShow   = "enter_stateShow"
	enterStateSettle = "enter_stateSettle"
)

//OnEnterStatusAny 进入任意阶段,都给房间内的所有session广播状态机状态改变
func (this *Desk) OnEnterStatusAny(e *fsm.Event) {
	logger.Log.Infof(fmt.Sprintf("状态机状态切换，deskId=%s,当前state=%s\n", this.deskID, this.deskFSM.Current()))
	broadcastInfo := &pb_bjl.BroadcastStatusChange{
		GameStatus: this.deskFSM.Current(),
	}
	pitaya.GroupBroadcast(context.Background(), "connector", this.deskGroupKey, "SeverBjl.ComponentRoom.SyncStatusChange", broadcastInfo)
}

func (this *Desk) OnEnterStatusReady(e *fsm.Event) {
	logger.Log.Infof(fmt.Sprintf("状态机状态切换，deskId=[%s],当前state=[%s]，event=%+v\n", this.deskID, this.deskFSM.Current(), e))
	//每5局洗一次牌
	if this.deskRound%5 == 0 {
		this.allPoker.Shuffle()
		logger.Log.Infof("OnEnterStatusReady, after shuffle, poker=%s", this.allPoker.String())
	}
	this.deskRound++
	//每局开始时，先清掉当局的下注流水
	this.roundBetInfo = make([]*pb_bjl.BetInfo, 0)
}

func (this *Desk) OnEnterStatusBet(e *fsm.Event) {
	logger.Log.Infof(fmt.Sprintf("状态机状态切换，deskId=[%s],当前state=[%s]，event=%+v\n", this.deskID, this.deskFSM.Current(), e))
}

func (this *Desk) OnEnterStatusSend(e *fsm.Event) {
	logger.Log.Infof(fmt.Sprintf("状态机状态切换，deskId=[%s],当前state=[%s]，[发牌]", this.deskID, this.deskFSM.Current()))
}

func (this *Desk) OnEnterStatusShow(e *fsm.Event) {
	this.curDeskPoker = NewDeskPoker(this)
	logger.Log.Infof(fmt.Sprintf("状态机状态切换，deskId=[%s],当前state=[%s]，开牌=%s", this.deskID, this.deskFSM.Current(), this.curDeskPoker.String()))
}

//OnEnterStatusSettle 结算
func (this *Desk) OnEnterStatusSettle(e *fsm.Event) {
	result := this.curDeskPoker.CalResult()
	this.historyResult = append(this.historyResult, result)
	if len(this.historyResult) > 30 {
		this.historyResult = this.historyResult[1:]
	}
	logger.Log.Infof(fmt.Sprintf("状态机状态切换，deskId=[%s],当前state=[%s]，结算=%+v\n", this.deskID, this.deskFSM.Current(), result))
}
