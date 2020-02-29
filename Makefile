# Make file

build:
	go build -o cinema-admin main.go

run:
	go run main.go

stop-services:
	docker-compose -f ./setup/docker-compose/db/postgreSQL.yml down

local-db:
	"======================== Setup DB (Postgres, Redis,...) ========================"

	docker-compose -f ./setup/docker-compose/db/postgreSQL.yml up -d
	docker cp ./setup/docker-compose/db/data.sql cinema-admin_postgres:/
	docker exec -u postgres cinema-admin_postgres psql cinema-admin user -f /data.sql
	
	"Ignore error warning. Setup db success"
	
setup-package:
	go get github.com/Masterminds/glide
	glide install
