## makefile中编译proto说明
* `--go_out=../`表示输出路径为当前文件夹的上级目录
* `-I=./proto`表示当前如果proto文件依赖了其它proto，那么从./proto目录中去找
* `./proto/*.proto`表示源文件在./proto目录下，*.proto仅将后缀为.proto文件生成proto

