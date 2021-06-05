package server_ddz

import (
	"github.com/topfreegames/pitaya/session"
	"server/pb/pb_enum"
)

type (
	Loser struct {
		uid   int64
		score uint
	}

	Player struct {
		Uid      int64            //用户ID
		avatar   string           //头像
		nickName string           //呢称
		sex      pb_enum.Sex      //性别
		session  *session.Session //玩家session
		desk     *Desk            //所在的桌子
		seat     int              //当前玩家在桌子的方位
		score    int              //当前玩家的分数
	}
)


