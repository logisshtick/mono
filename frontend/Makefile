NAME	= svlz

all: $(NAME)
$(NAME):
	pnpm i
	pnpm run build

c:
	docker build \
		-f ./build/Dockerfile-dev \
		-t $(NAME)-front-dev:latest \
		./

call: c
	docker run \
		--rm \
		--name $(NAME)-front-dev \
		--mount type=bind,source="$(PWD)/",target="/app" \
		-it \
		$(NAME)-front-dev:latest \
		"make"
