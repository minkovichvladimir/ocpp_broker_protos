.PHONY: all clean proto

all: proto

clean:
	rm -rf gen/go/*.pb.go

proto:
	protoc --go_out=./gen/go --go_opt=paths=source_relative \
		--go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative \
		proto/ocpp.proto 