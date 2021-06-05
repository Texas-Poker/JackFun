package poker

import (
	"math/rand"
	"time"
)

//pokerDzz 一组牌
type pokerDzz []int

//New 新建一副牌
func New(count int) pokerDzz {
	tiles := make(pokerDzz, count)

	for i := range tiles {
		tiles[i] = i
	}

	tiles.Shuffle()
	return tiles
}

//Shuffle 洗牌
func (this pokerDzz) Shuffle() {
	s := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range this {
		j := s.Intn(len(this))
		this[i], this[j] = this[j], this[i]
	}
}


