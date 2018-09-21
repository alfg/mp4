BINARY=mp4


build:
	go build

mp4info:
	go build -o ${BINARY} cmd/mp4info/mp4info.go

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -r ${BINARY}