PWD	= $(shell pwd)

all: up

gen-static:
	@if [ -d ./static/ ]; then \
		rm -rf static/; \
	fi

	cd ../frontend; \
		make call
	cp -r ../frontend/dist/ ../infra/static/

prepare:
	@if [ ! -f ./.env ]; then \
		touch .env; \
	fi
	
	@if [ ! -d ./logs/ ]; then \
		mkdir ./logs/; \
	fi

up: gen-static prepare
	docker compose up -d --build

down:
	docker compose down

clean:
	rm -r static
