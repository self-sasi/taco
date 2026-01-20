.PHONY: build run

build:
	go build -o bin/taco ./cmd/taco

run:
	go run ./cmd/taco $(filter-out $@,$(MAKECMDGOALS))

%:
	@: