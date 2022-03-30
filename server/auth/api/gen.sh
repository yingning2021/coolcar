protoc -I=. --go_out=paths=source_relative:gen/v1  auth.proto
protoc -I=. --go-grpc_out=paths=source_relative:gen/v1  auth.proto
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=auth.yaml:gen/v1 auth.proto