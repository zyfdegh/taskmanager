IMAGE_NAME=zyfdedh/taskmanager

default: help

help:
	@echo -e "Select a sub command \n"
	@echo -e "install-vendor: \n\t Install govendor"
	@echo -e "init-vendor: \n\t Init vendor/vendor.json"
	@echo -e "update-dep: \n\t Remove unused packages and add new packages in vendor/"
	@echo -e "get-dep: \n\t Synchronize packages from GOPATH or download online"
	@echo -e "build: \n\t Build Docker image"
	@echo -e "run: \n\t Run Docker container"
	@echo -e "push: \n\t Push image to DockerHub"
	@echo -e "fmt: \n\t Format source code with go fmt"
	@echo -e "help: \n\t Display this help"
	@echo -e "\n"
	@echo -e "See README.md for more."

install-vendor:
	go get github.com/kardianos/govendor

init-vendor:
	govendor init

update-dep:
	govendor remove +unused
	govendor add +external

get-dep:
	govendor sync

build:
	docker build -t ${IMAGE_NAME} .

run:
	docker run --rm -p 8082:8082 \
		${IMAGE_NAME}

push:
	docker push ${IMAGE_NAME}

fmt:
	go fmt $(go list ./... | grep -v /vendor/)
