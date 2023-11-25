
rm sql2java
rm sql2java.exe

# mac
go build sql2java.go

# windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build sql2java.go