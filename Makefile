NAME := gmsec
all: # 构建
	make clear
build:
	make clear
	./tools/gormt -o internal/model
	go build -o $(NAME) main.go 
	./$(NAME) debug
run:
	# make clear
	# ./tools/gormt -o internal/model
	go build -o $(NAME) main.go 
	./$(NAME) debug
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(NAME).exe main.go 
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(NAME) main.go 
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(NAME) main.go 
clear:
	- rm -rf  internal/model/* 
	- rm -rf err/ 
	- rm $(NAME)
	- rm $(NAME).exe
