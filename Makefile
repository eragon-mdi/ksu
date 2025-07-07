rebuild:
	docker compose down
	docker compose build
	docker compose up -d

restart:
	docker compose down
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

up:
	docker compose up -d

logs:
	docker logs ksu-app-task-1 | grep '^{.*}' | jq

lint:
	golangci-lint run ./...

#curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz
#sudo mv migrate /usr/local/bin
migrate_new:
	@if [ -z "$(name)" ]; then \
		echo "Error: укажи имя миграции через 'name=...'" && exit 1; \
	fi
	migrate create -ext sql -dir ./migrate $(name)

#clean:
#	sudo docker system prune -a --volumes