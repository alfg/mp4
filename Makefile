BINARY=mp4


.PHONY: test
test:
	go test ./...

build:
	go build

mp4info:
	go build -o ${BINARY} cmd/mp4info/mp4info.go


.PHONY: clean
clean:
	rm -r ${BINARY}