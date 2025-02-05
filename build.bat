@echo off
cd ./cli
go env -w GOOS=windows
go build -o ../docwiz.exe .