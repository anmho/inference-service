

default: api cli

.PHONY: gen clean

gen:
	@buf generate

clean:
	@rm -rf gen
	@rm -rf bin

api:
	@go build -o ./bin/api ./cmd/api

cli:
	@go build -o ./bin/client/ ./cmd/client
watch:
	@air --build.cmd "make api" --build.bin "./bin/api"

