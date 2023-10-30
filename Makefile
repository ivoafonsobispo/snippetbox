build:
	@go build -o bin/snippetbox

run: build
	@./bin/snippetbox