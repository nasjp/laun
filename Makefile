# CONST
GORUN=go run
GOBUILD=go build
GOCLEAN=go clean ./...
GOTEST=go test -cover ./...
DOCKER=docker exec -it
BINARY_NAME=laun
DOCKER_SERVICE_NAME=laun_app_1
MAIN=cmd/laun/

# VAR
var=

export GO111MODULE=on

all: build

build: test clean
	${GOBUILD} -o ${BINARY_NAME} ${MAIN}main.go

test:
	${GOTEST}

clean:
	${GOCLEAN}

docker:
	${DOCKER} ${DOCKER_SERVICE_NAME} ${GORUN} ${MAIN}main.go ${var}

clean-docker: clean-docer-comopse-images clean-docker-containers clean-docker-images;

clean-docker-containers:
	@echo ============================ start docker rm ===========================
	-@docker rm `docker container ls -qa`

clean-docker-images:
	@echo =========================== start docker rmi ===========================
	-@docker rmi -f `docker images -qa`

clean-docer-comopse-images:
	@echo ======================= start docker-compose down ======================
	-@docker-compose down --rmi all

# git remove marged branches without master
delete-branch:
	git branch --merged | grep -vE "^\*|master$$|develop$$" | xargs -I % git branch -d %

# debug with delve
dlv:
	cd ${MAIN} && dlv debug

# in docker container

in-docker-container:
	cd ${MAIN}; fresh