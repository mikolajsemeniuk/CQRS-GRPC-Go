# CQRS-GRPC-Go
```sh
go work init ./product-validation-service
go work use ./product-proxy-service
```
ElasticSearch get all indexes
```sh
curl --request GET --url 'http://localhost:9200/_cat/indices?v='
```
ElasticSearch create index
```sh
curl --request PUT --url http://localhost:9200/products
```
ElasticSearch remove index
```sh
curl --request DELETE --url http://localhost:9200/products
```
ElasticSearch get index schema
```sh
curl --request GET --url http://localhost:9200/products/_mapping
```
ElasticSearch get all documents
```sh
curl --request GET --url 'http://localhost:9200/products/_search?size=1308' | json_pp
```
ElasticSearch add/update document
```sh
curl --request PUT \
  --url http://localhost:9200/products/_doc/7e41edc1-8c5d-4902-af9d-b5dbea215a87 \
  --header 'Content-Type: application/json' \
  --data '{
	"id": "7e41edc1-8c5d-4902-af9d-b5dbea215a87",
	"name": "ps4",
	"dollars": 500,
	"cents": 0,
	"amount": 3,
	"is_imported": false,
	"created_at": "2022-04-23T18:25:43.511Z",
	"updated_at": "0000-00-00 00:00:00"
}'
```
ElasticSearch remove document
```sh
curl --request DELETE --url http://localhost:9200/products/_doc/7e41edc1-8c5d-4902-af9d-b5dbea215a87
```
What to do:
* changing to simple http router
* adding strict schema to elastic and see what happens