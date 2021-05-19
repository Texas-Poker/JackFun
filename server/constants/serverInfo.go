package constants

const (
	SvTypeConnector = "connector" //前端服务器
	SvTypeLobby     = "lobby"     //大厅服务器，用来处理登录、创建角色、角色名，用户信息等业务
	SvTypeGame      = "game"      //具体游戏业务服务器
	SvTypeWorld     = "world"     //世界服务器，用来广播消息，统计在线玩家等业务
)
