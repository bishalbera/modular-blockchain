build:
		go build -o ./bin/modular-blockchian
run: build
		./bin/modular-blockchian
test:
		go test ./...