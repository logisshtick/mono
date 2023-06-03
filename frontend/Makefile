NAME	= svlz

all: $(NAME)
$(NAME):
	pnpm i
	pnpm run build

lint:
	pnpm run lint

lint-fix:
	pnpm run lint-fix

prepare:
	@if [ ! -d ./dist ]; then \
		mkdir ./dist; \
	fi
	@if [ ! -d ./node_modules ]; then \
		mkdir ./node_modules; \
	fi
	@if [ ! -d ./.pnpm-store ]; then \
		mkdir ./.pnpm-store; \
	fi

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
