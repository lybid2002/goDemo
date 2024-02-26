set CGO_EABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o main-linux main.go
