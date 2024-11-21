

default: build

.PHONY: gen clean

gen:
	@buf generate

clean:
	@rm -rf gen
	@rm -rf bin

build:
	@go build -o ./bin/api ./cmd/api

watch:
	@air --build.cmd "make build" --build.bin "./bin/api"

