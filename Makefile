get_deps:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

generate:
	go generate ./...
	sqlc generate -f ./api/db/sqlc.yaml

migrate:
	migrate -path api/db/migration -database "postgres://localhost:5432/jackbox?sslmode=disable" up

prepare_frontend:
	npm install --prefix web/app build --force
	npm run --prefix web/app build
	cp -r ./web/app/dist ./cmd/user/frontend/dist

prepare_swagger:
	cp -r ./third_party/swagger ./cmd/user/swagger
	cp ./api/openapi/user/user.yml ./cmd/user/swagger

prepare_backend: get_deps generate prepare_swagger

prepare_with_migration: prepare_backend migrate

## prepare and move files to be embedded into the executable
## then clean up the moved files once built
prepare_and_build: prepare_frontend prepare_backend
	go build ./cmd/user/main.go
	rm -rf ./cmd/user/swagger
	rm -rf ./cmd/user/frontend

build_exec:
	go build ./cmd/user/main.go
	rm -rf ./cmd/user/swagger
	rm -rf ./cmd/user/frontend/dist

docker_build:
	docker build -t jackbox --file ./build/package/docker/Dockerfile .

docker_compose:
	docker compose -f deployments/docker-compose/docker-compose.yml up --build --force-recreate
