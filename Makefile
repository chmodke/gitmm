BUILD_NAME:=gitmm
BIN_DIR:=bin
BUILD_VERSION:=1.0.6
BUILD_DATE:=$(shell date '+%Y%m%d.%H%M%S')
SOURCE:=*.go
LDFLAGS:=-ldflags "-X '${BUILD_NAME}/cmd.VERSION=${BUILD_VERSION}-${BUILD_DATE}'"

all:clean build_win package

test: clean
	@echo "test"
	@go test

build_win: test
	@echo "build_win"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BIN_DIR}/${BUILD_NAME}.exe ${SOURCE}

build_unix: test
	@echo "build_unix"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BIN_DIR}/${BUILD_NAME} ${SOURCE}

install_win: clean build_win
	@echo "install_win"
	@cp ${BIN_DIR}/${BUILD_NAME}.exe /c/Windows/

install_unix: clean build_unix
	@echo "install_unix"
	@cp ${BIN_DIR}/${BUILD_NAME} /usr/bin

package :
	@echo "package"
	@cp README.md ${BIN_DIR}
	@cp repo.yaml ${BIN_DIR}
	- @[ -f ${BUILD_NAME}.tar.gz ] && rm ${BUILD_NAME}.tar.gz
	@cd ${BIN_DIR} && tar -czf ../${BUILD_NAME}.tar.gz *

clean:
	@echo "clean"
	@go clean
	- @rm -rf ${BIN_DIR}

.PHONY: all test build_win install_win build_unix install_unix package clean