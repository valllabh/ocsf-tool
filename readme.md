# Download OCSF Schema
schema.json is exported from below API
https://schema.ocsf.io/export/schema?profiles=cloud,container


# To install protoc
apt-get update
apt install -y protobuf-compiler

# To compile proto file
protoc ./*.proto --golang_out=./output
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
export PATH="$PATH:$(go env GOPATH)/bin"

go get -u google.golang.org/protobuf/cmd/protoc-gen-go