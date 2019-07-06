NAME=ctftime-bot
GOPKGNAME=github.com/kg0r0/ctftime-bot

build:
	cd $(GOPATH)/src/$(GOPKGNAME) && go build

clean:
	rm $(NAME)