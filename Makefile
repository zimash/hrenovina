all:
	go build -o ./bin/hrenovina

clean:
	rm -rf ./bin

lint:
	golangci-lint run -v
