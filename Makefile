.PHONY: all build run gotool clean help echo

BINARY="hertzTemplate"

## make - 格式化 Go 代码, 并编译生成二进制文件
all: gotool build

## make build - 编译 Go 代码, 生成二进制文件
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

## make run - 直接运行 Go 代码
run:
	@go run ./${BINARY}

## make gotool - 运行 Go 工具 'fmt' and 'vet'
gotool:
	go fmt ./
	go vet ./

## make clean - 移除二进制文件和 vim swap files
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help: Makefile
	@echo "\nUsage: make ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTINOS"
