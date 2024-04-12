package pb

/*
Get the Go gRPC tools by running

	$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

You need the tools to be in your path:

	$ export PATH="$PATH:$(go env GOPATH)/bin"

See more at https://grpc.io/docs/languages/go/quickstart/
*/

//to compile proto files without services
//protoc --go_out=. --go_opt=paths=source_relative ./tenant.proto

//to compile proto files with services
//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./tenant.proto
