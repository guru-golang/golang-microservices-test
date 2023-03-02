# Golang global paths

p			   := ""

export GOPATH  := $(shell pwd)/
export GOVERSION  := go1.18
export GORELACE  := rc
export PATH    := ${GOPATH}bin:${PATH}

# Docker repository
export NAME    := r.msp.yt/car-rent-platform/backend
export TAG     := 0.0.1 # $$(git describe --abbrev=0)
export IMG     := ${NAME}:${TAG}
export VERSION := ${NAME}:${GOVERSION}-${GORELACE}-${TAG}
export LATEST  := ${NAME}:latest
export DEVELOP := ${NAME}:dev


init:
	mkdir -p bin pkg src/${p}/src

mod-init:
	cd ./src/${p} && \
	go mod init ${m}

install:
	cd ./src/${p}/src && \
	go get -d -v

download:
	cd ./src/${p}/src && \
	go mod download

update: clean
	cd ./src/${p}/src && \
	go get -u -d -v

clean:
	cd ./src/${p}/src && \
	go clean --modcache

tidy:
	cd ./src/${p}/src && \
	go mod tidy

build:
	cd ./src/${p} && \
	go build -ldflags="-w -s" -buildvcs=false -o ./pkg/app ./src && \
	chmod +x ./pkg/app

run:
	cd ./src/${p} && \
	./pkg/app

enable-git-hooks:
	chmod u+x ./hooks/*
	git config core.hooksPath ${GOPATH}hooks
	git config advice.ignoredHook false

prepare: build-image push-image pull-image
build-image:
	@docker build --file="Dockerfile" --target="build" --build-arg="service=${p}" --tag="r.msp.yt/platforms/gitlab/devops/${p}:latest" ./
push-image:
	@docker push r.msp.yt/platforms/gitlab/devops/${p}:latest
pull-image:
	@docker --context=glab pull r.msp.yt/platforms/gitlab/devops/${p}:latest

stack:
	@docker --context=glab stack deploy glab-devops --compose-file stack.yml --with-registry-auth
swarm:
	make prepare p=backup
	make prepare p=restore
	make prepare p=transfer
	make prepare p=utility-cleaner
	make stack
