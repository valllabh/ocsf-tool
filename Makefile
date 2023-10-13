.PHONY: clean create-proto compile-proto run

clean:
	rm -rf ./output/*

create-proto:
	go run main.go

compile-proto:
	mkdir -p ./output/java
	mkdir -p ./output/golang
	find ./output/proto -type f -name "*.proto" | xargs protoc --proto_path=./output/proto --java_out=./output/java --go_opt=paths=source_relative --go_out=./output/golang

run: clean create-proto compile-proto
