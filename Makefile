build:
	go build -ldflags "-X github.com/lucaschain/beholder/cmd.Version=$(BEHOLDER_VERSION)" -o bin/beholder

run:
	go run main.go
