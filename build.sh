#!/bin/bash

cd ./cli || exit
go env -w GOOS=linux
go build -o ../docwiz .
