BINARY=mp4

.PHONY: all
all: mp4info

mp4info: go build -o ${BINARY} cmd/mp4info/mp4info.go

test: go test ./...

.PHONY: clean
clean: rm -r ${BINARY}