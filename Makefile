SHELL:=/bin/bash
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	DOCKER=sudo docker
endif
ifeq ($(UNAME_S),Darwin)
	DOCKER=docker
endif

install:
	go get ./...

build: install
	go build .

run: build
	./slack-message-redirector

test: install
	go test ./...

message:
	curl -X POST localhost:8000/slack -d "This is a testing body message"


message-remote:
	curl -X POST $$(kubectl get svc --selector='app=slack-message-redirector' -o jsonpath='{.items[0].status.loadBalancer.ingress[0].ip}')/slack -d "This is a testing body message"

DOCKER_IMAGE=chechiachang/slack-message-redirector

docker-build: test
	$(DOCKER) build -t $(DOCKER_IMAGE):latest -f Dockerfile .

docker: docker-build
	$(DOCKER) run --name slack-message-redirector -e SLACK_TOKEN=$${SLACK_TOKEN} -e SLACK_CHANNEL_ID=$${SLACK_CHANNEL_ID} -p 8000:8000 -d $(DOCKER_IMAGE)

docker-push: docker-build
	$(DOCKER) push $(DOCKER_IMAGE)
