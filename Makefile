bot:
	go run cmd/bot.go

build:
	go build -o bin/bot_arm64 cmd/bot.go

clean:
	go clean
	rm -rf bin
