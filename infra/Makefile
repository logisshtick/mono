PWD	= $(shell pwd)

all: up

gen-static:
	cd ../frontend; \
		make call; \
		mv ./static/ ../infra/static/; \
	cd $(PWD)

prepare:
	@if [ ! -f ./.env ]; then \
		touch .env; \
	fi
	
	@if [ ! -d ./logs/ ]; then \
		mkdir ./logs/; \
	fi

up: gen-static prepare
	docker compose up

down:
	docker compose down
