all: up

up:
	@if [ ! -f ./.env]; then \
		touch .env; \
	fi
	
	@if [ ! -d ./logs/ ]; then \
		mkdir ./logs/; \
	fi
	
	docker compose up

down:
	docker compose down
