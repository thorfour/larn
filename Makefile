.PHONY: game all clean

NAME="larn"

all: game
game:
	CGO_ENABLED=0 GOOS=linux go build -o ./$(NAME) ./cmd/
clean:
	rm $(NAME)
