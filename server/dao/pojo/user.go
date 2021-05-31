package pojo

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/topfreegames/pitaya/session"
	"server/dao/hashtree"
	"server/dao/redis_helper"
)

//User 执久层用户数据结构User(pojo_user)
type User struct {
	hashtree.Root
	Id       hashtree.Int64  `bson:"Id"`
	NickName hashtree.String `bson:"NickName"`
	Icon     hashtree.String `bson:"Icon"`
	Age      hashtree.Int64  `bson:"Age"`
	Sex      hashtree.Int64     `bson:"Sex"`
	Lv       hashtree.Int64  `bson:"Lv"`
	Gold     hashtree.Int64  `bson:"Gold"`
	Diamond  hashtree.Int64  `bson:"Diamond"`
	Token    hashtree.String `bson:"Token"`
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

func SetUserToSession(s *session.Session, u *User){
	s.Set("pojo_session", u)
}

// ManualSave 手动存redis
func (this *User) ManualSave() error {
	hash, err := this.Dump()
	if err != nil {
		return err
	}
	if len(hash) > 0 {
		p, err := redis_helper.GetAsyncPipeliner(redis_helper.User, this.Id.Get())
		if err != nil {
			return err
		}
		var cmds []redis.Cmder
		for key, value := range hash {
			cmds = append(cmds, redis.NewCmd("HSet", fmt.Sprintf("user[%v]", this.Id.Get()), key, value))
		}
		return p.DoCmds(cmds)
	}
	return nil
}
