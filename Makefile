build:
	GOARCH=arm64 go build -o bin/bot_arm64 .

clean:
	go clean
	rm -rf bin