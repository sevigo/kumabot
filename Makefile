BASEDIR=${CURDIR}
TMP=${BASEDIR}/tmp
VENDOR_TMP=${TMP}/vendor
LOCAL_BIN:=${TMP}/bin
GOBIN=${BASEDIR}/bin
PROJECT:=kumabot

VERSION ?= $(shell git describe --tags --abbrev=0)
GIT_REVISION = $(shell git rev-parse HEAD)

GO=go
GOFLAGS=
LOCAL_SERVICE_NAME=${GOBIN}/${PROJECT}
LDFLAGS += -X kumabot/pkg/version.Version=$(VERSION)
LDFLAGS += -X kumabot/pkg/version.Revision=$(GIT_REVISION)


run:
	@cd cmd/${PROJECT}; $(GO) run main.go inject_server.go inject_line_bot.go wire_gen.go

build:
	@cd cmd/${PROJECT} && $(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS) -w -s" -o $(LOCAL_SERVICE_NAME)

install:
	@echo "Installing ${PROJECT}"
	CGO_ENABLED=0 GOOS=linux ${GO} install -ldflags "$(LDFLAGS)" github.com/sevigo/kumabot/cmd/kumabot	

image-push:
	docker build -t sevigo/kumabot -f docker/Dockerfile .
	docker push sevigo/kumabot

wire:
	go get github.com/google/wire/cmd/wire

generate: wire
	go generate	./...

install-mockgen:
	GOPATH=${TMP} go get github.com/golang/mock/gomock
	GOPATH=${TMP} go install github.com/golang/mock/mockgen

mockgen: install-mockgen
	# grep "interface {" pkg/core/* | awk '{print $2}' | paste -sd "," - 
	${LOCAL_BIN}/mockgen -destination=mocks/mock_gen.go -package=mocks github.com/sevigo/kumabot/pkg/core KumaBots,KumaBot,LineBot

test:
	go test -timeout 10s -cover ./...
