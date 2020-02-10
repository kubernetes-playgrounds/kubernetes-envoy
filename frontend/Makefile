# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=envoy_frontend
BINARY_UNIX=$(BINARY_NAME)_unix

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker build -t guanw0826/kubernetes-envoy-frontend .

docker-push:
	docker push guanw0826/kubernetes-envoy-frontend

docker-run:
	docker run -d -p 8080:8080 guanw0826/kubernetes-envoy-frontend