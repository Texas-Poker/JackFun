package jcak_constants

const (
	SvTypeConnector = "sv_connector" //前端服务器
	SvTypeLobby     = "sv_lobby"     //大厅服务器，用来处理登录、创建角色、角色名，用户信息等业务
	SvTypeGame      = "sv_game"      //具体游戏业务服务器
	SvTypeWorld     = "sv_world"     //世界服务器，用来广播消息，统计在线玩家等业务
	SvTypeHttp      = "sv_http"
)
