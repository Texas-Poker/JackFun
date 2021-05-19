## 启动前端服务器
run_frontend_server:
	@go run main.go

## 启动后端服务器之大厅服务器
run_lobby_server:
	@go run main.go --port=3251 --type=lobby --frontend=false

## 编译protobuf
build_proto:
	@echo "step1:先清掉pb文件夹"
	@rm -rf pb
	@mkdir pb
	@echo "step2:再生成pb文件"
	@protoc --go_out=../ -I=../proto  ../proto/*.proto
	@echo "生成pb文件成功"

run_pitaya_cli:
	@cd ../pitaya-cli && ./pitaya-cli --file
