

# 1.编译为动态库
## 1.1 Linux
+ `go build -buildmode=c-shared -o grpc_auth.so -buildvcs=false`
## 1.2 Windows
+ `go build -buildmode=c-shared -o grpc_auth.dll`

# 2.测试
## 2.1 运行服务
+ `cd server`
+ `go run .\main.go`
## 2.2 编译为可执行文件
+ `cd test`
### 2.2.1 Linux
+ `gcc main.c -o main`
### 2.2.2 Windows
+ `gcc main.c -o main.exe`

# 3.执行
## 1.1 Linux
+ `./main`
## 1.2 Windows
+ `.\main.exe`