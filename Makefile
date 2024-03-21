BUILD_NAME:=gitmm
BIN_DIR:=bin
BUILD_MODE?=snapshot
BUILD_VERSION?=1.5.1
BUILD_DATE:=$(shell date '+%Y%m%d.%H%M%S')
SOURCE:=*.go

ifeq ($(BUILD_MODE),release)
SOFT_VERSION:=${BUILD_VERSION}
else
SOFT_VERSION:="${BUILD_VERSION}-${BUILD_DATE}"
endif

LDFLAGS:=-ldflags "-s -w -X 'github.com/chmodke/gitmm/cmd.VERSION=${SOFT_VERSION}'\
-X 'github.com/chmodke/gitmm/cmd.BuildId=${BUILD_DATE}'"

all:clean build_win build_linux package

prepare:
	@echo "prepare"
	- @go mod tidy

test: prepare clean
	@echo "test"
	- @go test ./...

build_win: prepare test
	@echo "build win"
	@export GO111MODULE=on \
		&& export CGO_ENABLED=0 \
		&& export GOOS=windows \
		&& export GOARCH=amd64 \
		&& go build ${LDFLAGS} -o ${BIN_DIR}/${BUILD_NAME}.exe ${SOURCE}
	- @upx ${BIN_DIR}/${BUILD_NAME}.exe

build_linux: prepare test
	@echo "build linux"
	@export GO111MODULE=on \
		&& export CGO_ENABLED=0 \
		&& export GOOS=linux \
		&& export GOARCH=amd64 \
		&& go build ${LDFLAGS} -o ${BIN_DIR}/${BUILD_NAME} ${SOURCE}
	- @upx ${BIN_DIR}/${BUILD_NAME}

install_win: clean prepare build_win
	@echo "install win"
	@cp ${BIN_DIR}/${BUILD_NAME}.exe /c/Windows/

install_linux: clean prepare build_linux
	@echo "install linux"
	@cp ${BIN_DIR}/${BUILD_NAME} /usr/bin

package :
	@echo "package"
	@cp README.md ${BIN_DIR}
	@cp repo.yaml ${BIN_DIR}
	- @find . -maxdepth 1 -type f -name "${BUILD_NAME}*.tar.gz" -exec rm -f {} \;
	@cd ${BIN_DIR} && tar -czf ../${BUILD_NAME}-${SOFT_VERSION}.tar.gz *

release :
	@echo "source release"
ifneq ($(BUILD_MODE),release)
	@echo "currently started in snapshot mode"
else
	@echo "${BUILD_VERSION} will be released"
	@git tag ${BUILD_VERSION} && git push --tag
endif

clean:
	@echo "clean"
	@go clean
	- @rm -rf ${BIN_DIR}
	- @find . -maxdepth 1 -type f -name "${BUILD_NAME}*.tar.gz" -exec rm -f {} \;

.PHONY: all test build_win install_win build_linux install_linux package clean