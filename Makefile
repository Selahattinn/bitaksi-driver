WHAT := bitaksi-driver

PWD ?= $(shell pwd)

VERSION   ?= $(shell git describe --tags)
REVISION  ?= $(shell git rev-parse HEAD)
BRANCH    ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILDUSER ?= $(shell id -un)
BUILDTIME ?= $(shell date '+%Y%m%d-%H:%M:%S')

.PHONY: build build-darwin-amd64 build-linux-amd64 build-windows-amd64 clean release test swagger

build: 
	for target in $(WHAT); do \
		go build -ldflags "-X github.com/Selahattinn/bitaksi-driver/pkg/version.Version=${VERSION} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.Revision=${REVISION} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.Branch=${BRANCH} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/$$target ./cmd/$$target; \
	done

build-darwin-amd64:
	for target in $(WHAT); do \
		CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -a -installsuffix cgo -ldflags "-X github.com/Selahattinn/bitaksi-driver/pkg/version.Version=${VERSION} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.Revision=${REVISION} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.Branch=${BRANCH} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/bitaksi-driver-${VERSION}-darwin-amd64/$$target ./cmd/$$target; \
	done

build-linux-amd64:
	for target in $(WHAT); do \
		CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -ldflags "-X github.com/Selahattinn/bitaksi-driver/pkg/version.Version=${VERSION} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.Revision=${REVISION} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.Branch=${BRANCH} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/bitaksi-driver-${VERSION}-linux-amd64/$$target ./cmd/$$target; \
	done

build-windows-amd64:
	for target in $(WHAT); do \
		CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -a -installsuffix cgo -ldflags "-X github.com/Selahattinn/bitaksi-driver/pkg/version.Version=${VERSION} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.Revision=${REVISION} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.Branch=${BRANCH} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/Selahattinn/bitaksi-driver/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/bitaksi-driver-${VERSION}-windows-amd64/$$target.exe ./cmd/$$target; \
	done

clean:
	rm -rf ./bin;

release: build-darwin-amd64 build-linux-amd64 build-windows-amd64
	cp ${PWD}/LICENSE ${PWD}/bin/bitaksi-driver-${VERSION}-darwin-amd64
	cp ${PWD}/LICENSE ${PWD}/bin/bitaksi-driver-${VERSION}-linux-amd64
	cp ${PWD}/LICENSE ${PWD}/bin/bitaksi-driver-${VERSION}-windows-amd64
	cp ${PWD}/config.yml ${PWD}/bin/bitaksi-driver-${VERSION}-darwin-amd64
	cp ${PWD}/config.yml ${PWD}/bin/bitaksi-driver-${VERSION}-linux-amd64
	cp ${PWD}/config.yml ${PWD}/bin/bitaksi-driver-${VERSION}-windows-amd64
	cd ${PWD}/bin; tar cfvz bitaksi-driver-${VERSION}-darwin-amd64.tar.gz ./bitaksi-driver-${VERSION}-darwin-amd64
	cd ${PWD}/bin; tar cfvz bitaksi-driver-${VERSION}-linux-amd64.tar.gz ./bitaksi-driver-${VERSION}-linux-amd64
	cd ${PWD}/bin; tar cfvz bitaksi-driver-${VERSION}-windows-amd64.tar.gz ./bitaksi-driver-${VERSION}-windows-amd64

test:
	go test ./...

swagger: 
	go get -u github.com/swaggo/swag/cmd/swag
	swag init --exclude ./cmd -g ./internal/api/api.go 