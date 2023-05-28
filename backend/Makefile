NAME	= svlz
VERSION	= 0.1
PORT	= 4433
LDFLAGS	= -s -w
# GCFLAGS	= -B -C

all: $(NAME)

$(NAME):
	go build -ldflags='$(LDFLAGS)' -gcflags=all='$(GCFLAGS)' -o $(NAME) ./cmd/$(NAME)/$(NAME).go 

container:
	docker build -f build/Dockerfile -t svlz-backend:$(VERSION) ./

up: container
	docker run -d --name svlz-backend --rm -p $(PORT):$(PORT) svlz-backend:$(VERSION)
