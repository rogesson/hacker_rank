build:
	docker build -t ruby-app .

run_bash: build
	docker run --name ruby-app --rm -i -t ruby-app bash

test: build
	docker run --rm ruby-app bash -c "rspec $(folder)/test.rb --color"

