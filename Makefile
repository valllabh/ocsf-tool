.PHONY: clean create-proto compile-proto run

install-protoc-linux:
	sudo apt-get update
	sudo apt-get install -y protobuf-compiler
	export PATH="$PATH:$(go env GOPATH)/bin"
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --version

clean-output-java:
	rm -rf ./output/java/*
clean-output-golang:
	rm -rf ./output/golang/*
clean-output-proto:
	rm -rf ./output/proto/*
clean-bin:
	rm -rf ./bin/*

clean-docs:
	rm -rf ./docs/*

clean-all: clean-output clean-bin clean-docs

build-docs: clean-docs
	go run doc-generator/doc-generator.go

build-project: clean-bin
	mkdir -p ./bin/
	go build -o ./bin/ocsf-tool

test-compile-proto: clean-output-java clean-output-golang
	mkdir -p ./output/java
	mkdir -p ./output/golang
	find ./output/proto -type f -name "*.proto" | xargs protoc --proto_path=./output/proto --java_out=./output/java --go_opt=paths=source_relative --go_out=./output/golang

test-create-proto: clean-output-proto
	./bin/ocsf-tool config extensions linux
	./bin/ocsf-tool config profiles cloud container
	./bin/ocsf-tool generate-proto file_activity

run: build-docs build-project test-create-proto test-compile-proto