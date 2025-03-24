.PHONY: all clean proto

all: proto

clean:
	rm -rf gen/go/ocpp_to_central/v1/*.pb.go
	rm -rf gen/go/ocpp_to_station/v1/*.pb.go
	rm -rf gen/go/mapping/v1/*.pb.go
	rm -rf gen/go/central_system/v1/*.pb.go

generate:
	@echo "Generating files..."
	protoc -I proto proto/ocpp_to_central/v1/ocpp_to_central.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	protoc -I proto proto/ocpp_to_station/v1/ocpp_to_station.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	protoc -I proto proto/mapping/v1/mapping.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	protoc -I proto proto/central_system/v1/central_system.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
