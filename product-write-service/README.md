
Create index
```sh
export ES_INDEX=products
go run product-write-service/commands/create-index/main.go
```
Drop index
```sh
export ES_INDEX=products
go run product-write-service/commands/drop-index/main.go
```
Seed Data
```sh
export ES_DATA_PATH=data.json
```
Generate Proto files
```sh
sudo protoc --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. \
    --go-grpc_opt=paths=source_relative \
    proto/*.proto
```
