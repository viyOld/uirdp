PROJECTNAME=$(shell basename "$(PWD)")

# Go переменные.
#GOBASE=$(shell pwd)
#GOPATH=$(GOBASE)/vendor:$(GOBASE):/home/azer/code/golang  #Вы можете удалить или изменить путь после двоеточия.
#GOBIN=$(GOBASE)/bin
#GOFILES=$(wildcard *.go)

# Перенаправление вывода ошибок в файл, чтобы мы показывать его в режиме разработки.
#STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID-файл будет хранить идентификатор процесса, когда он работает в режиме разработки
#PID=/tmp/.$(PROJECTNAME)-api-server.pid

# Make пишет работу в консоль Linux. Сделаем его silent.
#MAKEFLAGS += --silent

build:
	go build main.go
#    go build -o bin/main main.go
#
run:
	go run main.go helpPage.go
compile:
# 32-Bit Systems
# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386 main.go
# Linux
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
# 64-Bit
# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/main-freebsd-amd64 main.go
# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 main.go
# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go
