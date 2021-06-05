package server_bjl

import "server/pb/pb_bjl"

type deskHistory struct {
	ID      uint32
	BetInfo []*pb_bjl.BetInfo //桌子上当前局的下注流水
	Pokers  *deskPoker        //桌子上发的牌
	result  *pb_bjl.Result    //该局的结果
}

func NewHistory(id uint32, betInfo []*pb_bjl.BetInfo, pokers *deskPoker, result *pb_bjl.Result) *deskHistory {
	return &deskHistory{
		ID:      id,
		BetInfo: betInfo,
		Pokers:  pokers,
		result:  result,
	}
}

