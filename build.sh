#!/bin/bash
cd ./cli
go env -w GOOS=linux
go build -o ../docwiz .