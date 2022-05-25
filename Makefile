ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif

.DEFAULT_GOAL := docker-push-image

# local dev
test:
	go test ./...
.PHONY:test

build: test
	go build main.go
.PHONY:build

run-with-env: build
	docker run --rm -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d postgres:latest
	$$env:APP_DB_USERNAME='postgres'; $$env:APP_DB_PASSWORD='postgres'; $$env:APP_DB_NAME='test'; ./main
.PHONY:run-with-env

# clear local dev
clear:
	rm main.exe
.PHONY:clear

# docker
docker-build-image:
	docker build -t stakkato95/lambda-ingress:latest . -f Dockerfile
.PHONY:docker-build-image

docker-push-image: docker-build-image
	docker push stakkato95/lambda-ingress:latest
.PHONY:docker-push-image

# docker-run-tmp-container:
# 	docker run --rm -p 8080:8080 -d stakkato95/lambda-ingress
# .PHONY:docker-local-container
