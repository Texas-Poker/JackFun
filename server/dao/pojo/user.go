package pojo

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/topfreegames/pitaya/session"
	t "server/dao/hashtree"
	"server/dao/redis_helper"
	"server/pb/pb_enum"
)

//User 执久层用户数据结构User(pojo_user)
type (
	User struct {
		t.Root
		UID  t.Int64
		Char CharInfo
	}

	CharInfo struct {
		NickName t.String
		Icon     t.String
		Age      t.Uint32
		Sex      t.Uint32
		Lv       t.Uint32
		Gold     t.Uint32
		Diamond  t.Uint32
	}
)

func NewUser(uid int64, account string) *User {
	u := new(User)
	u.Load(u,nil)
	u.UID.Set(uid)
	u.Char.Diamond.Set(1000)
	u.Char.NickName.Set(fmt.Sprintf("玩家%s", account))
	u.Char.Sex.Set(uint32(pb_enum.Sex_Unknow))
	u.Char.Gold.Set(2000)
	u.Char.Lv.Set(1)
	u.ManualSave()
	return u
}

func  UserKey(uid int64) string {
	return fmt.Sprintf("user[%d]", uid)
}

func GetUserFromSession(s *session.Session) *User {
	if s == nil {
		return nil
	}
	v := s.Value("pojo_session")
	if v == nil {
		return nil
	}
	return v.(*User)
}

func SetUserToSession(s *session.Session, u *User) {
	s.Set("pojo_session", u)
}

// ManualSave 手动存redis
func (this *User) ManualSave() error {
	hash, err := this.Dump()
	if err != nil {
		return err
	}
	if len(hash) > 0 {
		p, err := redis_helper.GetAsyncPipeliner(redis_helper.User, this.UID.Get())
		if err != nil {
			return err
		}
		var cmds []redis.Cmder
		for key, value := range hash {
			cmds = append(cmds, redis.NewCmd("HSet", UserKey(this.UID.Get()), key, value))
		}
		return p.DoCmds(cmds)
	}
	return nil
}
