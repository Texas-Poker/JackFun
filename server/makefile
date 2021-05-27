## 启动前端服务器
run_server_frontend:
	@go run main.go --port=3250 --type=ServerConnector --frontend=true

## 启动后端服务器之大厅服务器
run_server_lobby:
	@go run main.go --port=3251 --type=ServerLobby --frontend=false

run_server_http:
	@go run main.go --port=8088 --type=ServerHttp --frontend=false

## 编译protobuf
build_proto_go:
	@echo "step1:先清掉pb文件夹"
	@rm -rf pb
	@mkdir pb
	@echo "step2:再生成pb文件"
	@protoc --go_out=../ -I=../proto  ../proto/*.proto
	@echo "生成pb文件成功"

## 安装pitaya-cli,如果安装失败，可下载源码自己编译后放到go path下的bin目录
install_pitaya_cli:
	@go get -u github.com/topfreegames/pitaya-cli

## 以pb解析形式启动pitaya-cli
run_pitaya_cli_pb:
	@#cd ..
	@cd ../pitaya-cli
	@go run main.go --docs=6

## 以json解析形式启动pitaya-cli
run_pitaya_cli_json:
	@pitaya-cli

## 为客户端生成protobuf
build_proto_cs:
	@echo "step1:先清掉pb文件夹"
	@rm -rf ../client/Assets/Pb
	@mkdir ../client/Assets/Pb
	@echo "step2:再生成pb文件"
	@protoc --csharp_out=../client/Assets/Pb -I=../proto  ../proto/*.proto
	@echo "生成pb文件成功"

build_proto_all:
	@make build_proto_cs
	@make build_proto_go