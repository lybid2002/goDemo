set CGO_EABLED=0
set GOOS=darwin
set GOARCH=amd64
go build -o main-mac main.go
# echo 执行，直接命令行：./
# 如果提示无权限： chmod 777 main
