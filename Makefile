build:
	@go build -o bin/snippetbox ./cmd/web

run: build
	@./bin/snippetbox