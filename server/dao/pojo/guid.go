package pojo

type GuidType int

const (
	GuidTypeUser      GuidType = 1 //玩家guid类型
	GuidTypeMail      GuidType = 2 //邮件guid类型
	GuidTypeItem      GuidType = 3 //各种物品item类型
	GuidTypeRedPacket GuidType = 4 //红包的guid类型
)

type GuidData struct {

	Id    int
	Perch int
	Low   int
	Des   string
}

func (g *GuidData) TableName() string {
	return "table_guid_data"
}
