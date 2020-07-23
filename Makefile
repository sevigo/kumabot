BASEDIR=${CURDIR}
TMP=${BASEDIR}/tmp
VENDOR_TMP=${TMP}/vendor
LOCAL_BIN:=${TMP}/bin
PROJECT:=kumabot

GO = /snap/bin/go

run:
	cd cmd/${PROJECT}; go run main.go inject_server.go inject_line_bot.go wire_gen.go

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
