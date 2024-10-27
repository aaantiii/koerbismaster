bot:
	go run cmd/bot.go

build:
	go build -o bin/bot cmd/bot.go

clean:
	go clean
	rm -rf bin
