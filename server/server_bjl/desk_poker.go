package server_bjl

import (
	"fmt"
	"server/pb/pb_bjl"
	"server/server_bjl/poker"
)

type deskPoker struct {
	desk *Desk

	xian1 *poker.Poker
	xian2 *poker.Poker
	xian3 *poker.Poker

	zhuang1 *poker.Poker
	zhuang2 *poker.Poker
	zhuang3 *poker.Poker
}

func NewDeskPoker(desk *Desk) *deskPoker {
	target := &deskPoker{
		desk: desk,

		xian1: desk.allPoker.FaPai(),
		xian2: desk.allPoker.FaPai(),

		zhuang1: desk.allPoker.FaPai(),
		zhuang2: desk.allPoker.FaPai(),
	}

	isXianBuPai, xian3 := target.IsXianBuPai()
	if isXianBuPai {
		target.xian3 = xian3
	}

	isZhuangBuPai, zhuang3 := target.IsZhuangBuPai(isXianBuPai)
	if isZhuangBuPai {
		target.zhuang3 = zhuang3
	}
	return target
}

func (this *deskPoker) TotalXianScore() int {
	totalScore := 0
	totalScore += this.xian1.Score + this.xian2.Score
	if this.xian3 != nil {
		totalScore += this.xian3.Score
	}
	totalScore = totalScore % 10
	return totalScore
}

func (this *deskPoker) TotalZhuangScore() int {
	totalScore := 0
	totalScore += this.zhuang1.Score + this.zhuang2.Score
	if this.zhuang3 != nil {
		totalScore += this.zhuang3.Score
	}
	totalScore = totalScore % 10
	return totalScore
}

func (this *deskPoker) IsXianBuPai() (bool, *poker.Poker) {
	if this.TotalXianScore() <= 7 {
		return true, this.desk.allPoker.FaPai()
	}
	return false, nil
}

func (this *deskPoker) IsZhuangBuPai(isXianBuPai bool) (bool, *poker.Poker) {
	zhuangScore := 0
	zhuangScore += this.zhuang1.Score + this.zhuang2.Score
	zhuangScore = zhuangScore % 10

	switch zhuangScore {
	case 0, 1, 2:
		return true, this.desk.allPoker.FaPai()
	case 3:
		if isXianBuPai && this.xian3.Score == 8 {
			return false, nil
		}
		return true, this.desk.allPoker.FaPai()
	case 4:
		if isXianBuPai && (this.xian3.Score == 0 || this.xian3.Score == 1 || this.xian3.Score == 8 || this.xian3.Score == 9) {
			return false, nil
		}
		return true, this.desk.allPoker.FaPai()
	case 5:
		if isXianBuPai && (this.xian3.Score == 0 || this.xian3.Score == 1 || this.xian3.Score == 2 || this.xian3.Score == 3 || this.xian3.Score == 8 || this.xian3.Score == 9) {
			return false, nil
		}
		return true, this.desk.allPoker.FaPai()
	case 6:
		if isXianBuPai && (this.xian3.Score == 6 || this.xian3.Score == 7) {
			return false, nil
		}
		return true, this.desk.allPoker.FaPai()
	default:
		return false, nil
	}
}

func (this *deskPoker) String() string {
	str := ""
	str += fmt.Sprintf("闲1=%s,闲2=%s", this.xian1, this.xian2)
	if this.xian3 != nil {
		str += fmt.Sprintf(",闲3=%s", this.xian3)
	}

	str += fmt.Sprintf(" ,庄1=%s,庄2=%s", this.zhuang1, this.zhuang2)
	if this.zhuang3 != nil {
		str += fmt.Sprintf(",庄3=%s", this.zhuang3)
	}

	return str
}

func (this *deskPoker)CalResult() *pb_bjl.Result {
	result := &pb_bjl.Result{}
	if this.TotalXianScore() > this.TotalZhuangScore() {
		result.WinType = pb_bjl.EnumWinType_Xian
	} else if this.TotalZhuangScore()  > this.TotalXianScore() {
		result.WinType = pb_bjl.EnumWinType_Zhuang
	} else {
		result.WinType = pb_bjl.EnumWinType_He
	}

	if this.xian1.Point == this.xian2.Point {
		result.IsXianDui = true
	} else {
		result.IsXianDui = false
	}

	if this.zhuang1.Point == this.zhuang2.Point {
		result.IsZhuangDui = true
	} else {
		result.IsZhuangDui = false
	}

	return result
}
