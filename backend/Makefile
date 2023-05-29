NAME	= svlz
VERSION	= 0.1
PORT	= 4433
LDFLAGS	= -s -w
# GCFLAGS	= -B -C

all: $(NAME)

gen:
	go generate ./...

$(NAME): gen
	go build \
		-ldflags='$(LDFLAGS)' \
		-gcflags=all='$(GCFLAGS)' \
		-o $(NAME) \
		./cmd/$(NAME)/$(NAME).go 

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
		--mount type=bind,source="./",target="/app" \
		-it \
		$(NAME)-dev:$(VERSION) \
		"make fmt"

c:
	docker build \
		-f ./build/Dockerfile-prod \
		-t $(NAME)-backend:$(VERSION) \
		./

up: c
	docker run \
		-d \
		--rm \
		--name $(NAME)-backend \
		-p $(PORT):4433 \
		$(NAME)-backend:$(VERSION)
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
