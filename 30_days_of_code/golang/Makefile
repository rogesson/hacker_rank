build:
	docker build -t go-app .

run:
	docker run --rm go-app bash -c "go run $(file) < in.txt"

run_bash:
	docker run --name go-app --rm -i -t go-app bash

#[todo] add syntax analyzer
#lint:
#	docker run --rm go-app ???

