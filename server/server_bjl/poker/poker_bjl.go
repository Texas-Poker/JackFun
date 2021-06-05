package poker

import (
	"fmt"
	"math/rand"

	"time"
)

//PokerBjl 一组牌
type PokerBjl struct {
	index    uint
	allPoker []*Poker
}

//NewPokerBjl 新建一副牌
func NewPokerBjl(count int) *PokerBjl {
	pokerBjl := &PokerBjl{
		index:    0,
	}

	pokerBjl.allPoker = make([]*Poker, 0)
	for i := 0; i < count; i++ {
		p := NewPoker(i)
		pokerBjl.allPoker = append(pokerBjl.allPoker, p)
	}

	pokerBjl.Shuffle()
	return pokerBjl
}

//Shuffle 洗牌
func (this *PokerBjl) Shuffle() {
	s := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range this.allPoker {
		j := s.Intn(len(this.allPoker))
		this.allPoker[i], this.allPoker[j] = this.allPoker[j], this.allPoker[i]
	}
}

//FaPai 发牌
func (this *PokerBjl) FaPai() *Poker {
	target := this.allPoker[this.index]
	this.index+=1
	return target
}

func (this *PokerBjl) String() string {
	str := ""
	for i, poker := range this.allPoker {
		str += fmt.Sprintf("[%d]=%s ", i, poker.String())
	}
	return str
}
