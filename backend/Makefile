NAME = svlz
LDFLAGS = -s -w
# GCFLAGS = -B -C

all: $(NAME)

$(NAME):
	go build -ldflags='$(LDFLAGS)' -gcflags=all='$(GCFLAGS)' cmd/$(NAME)/$(NAME).go -o $(NAME)
