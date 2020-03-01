# Make file

build:
	go build -o cinema-admin main.go

run:
	go run main.go

stop-services:
	docker-compose -f ./setup/docker-compose/postgreSQL.yml down

local-db:
	docker-compose -f ./setup/docker-compose/postgreSQL.yml up -d
	docker cp ./setup/docker-compose/data.sql cinema-admin_postgres:/
	docker exec -u postgres cinema-admin_postgres psql cinema-admin user -f /data.sql

setup-package:
	mkdir logs
	go get github.com/Masterminds/glide
	glide install

test:
	go test -v ./...

docker-image:
	make build 