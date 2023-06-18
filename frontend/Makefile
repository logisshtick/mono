NAME	= svlz

all: build

prepare:
	@if [ ! -d ./node_modules ]; then \
		mkdir ./node_modules; \
	fi
	@if [ ! -d ./.pnpm-store ]; then \
		mkdir ./.pnpm-store; \
	fi
	@if [ ! -d ./dist ]; then \
		mkdir ./dist; \
	fi

dep-install:
	pnpm i

build: dep-install
	pnpm run build

lint:
	pnpm run lint

lint-fix:
	pnpm run lint-fix

c:
	docker build \
		-f ./build/Dockerfile-dev \
		-t $(NAME)-front-dev:latest \
		./

call: prepare c
	docker run \
		--rm \
		--name $(NAME)-front-dev \
		--mount type=bind,source="$(PWD)/",target="/app" \
		-it \
		$(NAME)-front-dev:latest \
		"make"

clean:
	rm -rf ./dist
	rm -rf ./node_modules
	rm -rf ./.pnpm-store
