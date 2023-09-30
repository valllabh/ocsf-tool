.PHONY: create-proto compile-proto run

create-proto:
	go run main.go

compile-proto:
	mkdir -p ./output/java
	protoc ./*.proto --java_out=./output/java

run: create-proto compile-proto
