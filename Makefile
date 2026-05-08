
tidy:
	go mod tidy

format:
	go fmt ./...

lint:
	go vet ./...

check:
	staticcheck ./...

import:
	find . -name '*.go' -exec goimports -w {} +

unitt:
	go test `go list ./... | grep -v /store` -v -coverprofile=coverage.out
	# go test -v ./... -coverprofile=coverage.out
	# go tool cover -html=coverage.out -o coverage.html

dbt:
	go test `go list ./... | grep /store` -v -coverprofile=coverage.out
	# go tool cover -html=coverage.out -o coverage.html

singlet:
	go test -v $(file)

migration:
	migrate create -dir ./pkg/${model}/schema -ext sql $(name)

dockerup:
	docker compose -f build/compose.yaml up -d

docker:
	docker compose -f build/compose.yaml $(cmd)

postgres:
	docker compose -f build/compose.yaml exec postgres psql -d postgres -U postgres

migrate:
	mkdir -p tmp/migrations && cp pkg/*/schema/* tmp/migrations/ && migrate -source file://tmp/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

dockerlog:
	docker compose -f build/compose.yaml logs app

dockerbuild:
	docker compose -f build/compose.yaml rm app && docker compose -f build/compose.yaml build app

runt:
	PORT=${port} runn run test/index.yml --scopes run:exec
