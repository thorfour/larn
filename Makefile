.PHONY: all clean

NAME="larn"

all: clean larn
larn:
	CGO_ENABLED=0 GOOS=linux go build -o ./$(NAME) ./cmd/larn
clean:
	@[ -f $(NAME) ] && rm $(NAME) || true
debug:
	CGO_ENABLED=0 GOOS=linux go build -tags DEBUG -o ./$(NAME) ./cmd/larn
