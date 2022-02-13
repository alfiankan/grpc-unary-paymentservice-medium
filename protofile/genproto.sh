mkdir -p $1
protoc --go_out=../$1 --go_opt=paths=source_relative \
    --go-grpc_out=../$1 --go-grpc_opt=paths=source_relative \
    $2
