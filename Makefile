

.PHONY: gen

gen:
	@buf generate

clean:
	@rm -rf gen
