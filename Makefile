FILE_PATH := ./224/a.go

build:
	go build -buildmode=exe -o ./a.out $(FILE_PATH)

run:
	./a.out

build-run:
	@make build
	@make run
