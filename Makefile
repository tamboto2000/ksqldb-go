gen-parser:
	java -jar "./tools/bin/antlr-4.13.0-complete.jar" -Dlanguage=Go -o parser antlr/v0-29-x/SqlBase.g4

lint:
	golangci-lint run

start-ksqldb:
	docker-compose up -d

stop-ksqldb:
	docker-compose down

restart-ksqldb: stop-ksqldb start-ksqldb

ksql-cli:
	docker exec -it ksqldb-cli ksql http://primary-ksqldb-server:8088			

test-cov:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out

test-cov-view:
	go tool cover -html="coverage.out"