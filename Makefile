build:
	go build \
		-ldflags "-X github.com/lucaschain/beholder/cmd.Version=$(BEHOLDER_VERSION)" \
		-o bin/beholder



run:
	go run main.go

test:
	go test -v ./...


.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest


.PHONY: check-coverage
check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml
