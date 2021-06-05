package poker

var SuitName = []string{"桃", "心", "梅", "方"}
var MaxTileIndex = 52

type Poker struct {
	Id    int
	Suit  int //花色
	Point int //点数
	Score int //分值
}

// 花色
const (
	flowerNIL      int = iota // 留空
	flowerHEITAO              // 黑桃(小王)
	flowerHONGTAO             // 红桃(大王)
	flowerMEIHUA              // 梅花
	flowerFANGKUAI            // 方块
)

// 点数
const (
	cardPointNIL int = iota // 留空
	cardPointA
	cardPoint2
	cardPoint3
	cardPoint4
	cardPoint5
	cardPoint6
	cardPoint7
	cardPoint8
	cardPoint9
	cardPointT
	cardPointJ
	cardPointQ
	cardPointK
	cardPointX // 小王
	cardPointY // 大王
)

// NewPoker 新建卡牌
func NewPoker(id int) *Poker {
	p := &Poker{
		Id:   id,
		Suit: toFlower(id),
	}
	p.Point, p.Score = toPointAndScore(id)
	return p
}

//String 转换成字符串
func (this *Poker) String() string {
	strResult := ""
	// 花色
	switch this.Suit {
	case flowerHEITAO:
		{
			strResult = "♠"
		}
	case flowerHONGTAO:
		{
			strResult = "♥"
		}
	case flowerMEIHUA:
		{
			strResult = "♣"
		}
	case flowerFANGKUAI:
		{
			strResult = "♦"
		}
	}

	// 点数
	switch this.Point {
	case cardPoint3:
		{
			strResult = strResult + "3"
		}
	case cardPoint4:
		{
			strResult = strResult + "4"
		}
	case cardPoint5:
		{
			strResult = strResult + "5"
		}
	case cardPoint6:
		{
			strResult = strResult + "6"
		}
	case cardPoint7:
		{
			strResult = strResult + "7"
		}
	case cardPoint8:
		{
			strResult = strResult + "8"
		}
	case cardPoint9:
		{
			strResult = strResult + "9"
		}
	case cardPointT:
		{
			strResult = strResult + "T"
		}
	case cardPointJ:
		{
			strResult = strResult + "J"
		}
	case cardPointQ:
		{
			strResult = strResult + "Q"
		}
	case cardPointK:
		{
			strResult = strResult + "K"
		}
	case cardPointA:
		{
			strResult = strResult + "A"
		}
	case cardPoint2:
		{
			strResult = strResult + "2"
		}
	case cardPointX:
		{
			strResult = "小王"
		}
	case cardPointY:
		{
			strResult = "大王"
		}
	}
	return strResult
}

// 从牌值获取花色
func toFlower(id int) int {
	if id <= 0 || id > 54 {
		return flowerNIL
	}

	return ((id - 1) / 13) + 1
}

// 从牌值获取点数
func toPointAndScore(id int) (point int, score int) {
	if id <= 0 {
		return cardPointNIL, 0
	}
	if id == 53 {
		return cardPointX, 0 // 小王
	}
	if id == 54 {
		return cardPointY, 0 // 大王
	}
	point = (id-1) % 13+1
	if point == cardPointT || point == cardPointJ || point == cardPointQ || point == cardPointK {
		score = 0
	} else {
		score = point
	}
	return point, score
}
