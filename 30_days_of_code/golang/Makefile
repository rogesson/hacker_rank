build:
	docker build -t go-app .

run: build
	docker run --rm go-app bash -c "go run $(file) < in.txt"

run_bash: build
	docker run --name go-app --rm -i -t go-app bash

test: build
	docker run --rm go-app bash -c "go test -v $(folder)/main_test.go $(folder)/main.go"

#[todo] add syntax analyzer
#lint:
#	docker run --rm go-app ???

