NAME	= svlz
VERSION	= 0.1
PORT	= 4433
LOGFILENAME	= $(shell date '+%Y-%m-%d %H-%M-%S').log
PWD		= $(shell pwd)
LDFLAGS	= -s -w
# GCFLAGS	= -B -C
CGO_ENABLED	= 0

all: $(NAME)

gen:
	go generate ./...

$(NAME): gen
	go build \
		-ldflags='$(LDFLAGS)' \
		-gcflags=all='$(GCFLAGS)' \
		-o $(NAME) \
		./cmd/$(NAME)/$(NAME).go 
clean:
	rm ./$(NAME)

fmt:
	go fmt ./...
cfmt:
	docker build \
		-f ./build/Dockerfile-dev \
		-t $(NAME)-dev:$(VERSION) \
		./
	docker run \
		--rm \
		--name $(NAME)-dev \
		--mount type=bind,source="$(PWD)/",target="/app" \
		-it \
		$(NAME)-dev:$(VERSION) \
		"make fmt"

c:
	docker build \
		-f ./build/Dockerfile-prod \
		-t $(NAME)-backend:$(VERSION) \
		./

up: c
	@if [ ! -d ./logs/ ]; then \
		mkdir ./logs/; \
	fi
	docker run \
		-d \
		--rm \
		--name $(NAME)-backend \
		--mount type=bind,source="$(PWD)/logs/",target="/app/logs" \
		-p $(PORT):$(PORT) \
		$(NAME)-backend:$(VERSION) \
		-port $(PORT) \
		-log-file "./logs/$(LOGFILENAME)"
	@echo
	@echo ----------------------------------
	@echo Upped on port: $(PORT)
	@echo Connect: http:/127.0.0.1:$(PORT)/
	@echo ----------------------------------
	@echo for stop exec: make down

down:
	docker stop $(NAME)-backend

# for all how need it)
sync:
	git pull


