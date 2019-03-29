install:
	go get ..

build: install
	go build .

run: build
	./slack-message-redirector

test:
	go test ./...

message:
	curl -X POST localhost:8000/slack -d "This is a testing body message"

DOCKER_IMAGE=chechiachang/slack-message-redirector

docker-build: test
	docker build -t $(DOCKER_IMAGE):latest -f Dockerfile .

docker: docker-build
	docker run --name slack-message-redirector -p 8000:8000 -d $(DOCKER_IMAGE)
