protoc \
--go_out=proto_common \
--go_opt=paths=source_relative \
common.proto

protoc \
--go_out=proto_book --go-grpc_out=proto_book \
--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
book.proto
