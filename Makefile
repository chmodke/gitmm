BUILD_NAME:=gitmm
BIN_DIR:=bin
BUILD_MODE?=snapshot
BUILD_VERSION?=1.1.0
BUILD_DATE:=$(shell date '+%Y%m%d.%H%M%S')
SOURCE:=*.go

ifeq ($(BUILD_MODE),release)
SOFT_VERSION:=${BUILD_VERSION}
else
SOFT_VERSION:="${BUILD_VERSION}-${BUILD_DATE}"
endif

LDFLAGS:=-ldflags "-X '${BUILD_NAME}/cmd.VERSION=${SOFT_VERSION}'"

all:clean build_win package

test: clean
	@echo "test"
	- @go test ./...

build_win: test
	@echo "build win"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BIN_DIR}/${BUILD_NAME}.exe ${SOURCE}

build_linux: test
	@echo "build linux"
	@GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BIN_DIR}/${BUILD_NAME} ${SOURCE}

install_win: clean build_win
	@echo "install win"
	@cp ${BIN_DIR}/${BUILD_NAME}.exe /c/Windows/

install_linux: clean build_linux
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
	@git tag ${BUILD_VERSION}
	@git push --tag
endif

clean:
	@echo "clean"
	@go clean
	- @rm -rf ${BIN_DIR}
	- @find . -maxdepth 1 -type f -name "${BUILD_NAME}*.tar.gz" -exec rm -f {} \;

.PHONY: all test build_win install_win build_linux install_linux package clean