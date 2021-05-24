## protobuf编译
*目录结构图如下
``` graph
jack_fun
│
└───client
│
└───proto
│    └─── .proto                --这里是各种proto源文件
└───server
    └─── pb                     --这里是存放编译后的proto对应的go文件
    └─── Makefile               --命令集
```
* 图中`proto`文件夹用来存放`client`和`server`共用的proto文件
* 在`server/pb`文件夹用来存放编译后的proto对应的go文件
* 使用Makefile中命令`make build_proto`用来编译proto
* 上述命令将执行`protoc --go_out=../ -I=../proto  ../proto/*.proto`  
> `--go_out=../`表示输出路径为当前文件夹的上级目录（即图中的jack_fun目录）
> `-I=../proto`表示当前如果proto文件依赖了其它proto，那么从../proto目录中去找（即图中的jack_fun/proto目录）
> `../proto/*.proto`表示源文件在../proto目录下（即图中的jack_fun/proto目录），*.proto仅将后缀为.proto文件

