# Make file

build:
	go build -o cinema-admin main.go

run:
	go run main.go

stop-services:
	docker-compose -f ./setup/docker-compose/postgreSQL.yml down
	docker-compose -f ./setup/docker-compose/app.yml down
	docker volume rm docker-compose_database-data

local-db:
	docker-compose -f ./setup/docker-compose/postgreSQL.yml up -d
	@while ! docker exec cinema-admin_postgres pg_isready -h localhost -p 5432 > /dev/null; do \
		sleep 1; \
	done
	docker cp ./setup/docker-compose/data.sql cinema-admin_postgres:/
	docker exec -u postgres cinema-admin_postgres psql cinema-admin user -f /data.sql

app:
	docker-compose -f ./setup/docker-compose/app.yml up -d

setup-package:
	mkdir logs
	go get github.com/Masterminds/glide
	glide install

test:
	go test -v ./...

docker-image:
	make build 