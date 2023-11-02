.PHONY: clean create-proto compile-proto run

clean:
	rm -rf ./docs/*
	rm -rf ./output/*
	rm -rf ./bin/*

docs:
	go run doc-generator/doc-generator.go

build:
	mkdir -p ./bin/
	go build -o ./bin/ocsf-tool

compile-proto:
	mkdir -p ./output/java
	mkdir -p ./output/golang
	find ./output/proto -type f -name "*.proto" | xargs protoc --proto_path=./output/proto --java_out=./output/java --go_opt=paths=source_relative --go_out=./output/golang

test-run:
	./bin/ocsf-tool generate proto file_activity security_finding

run: clean docs build test-run