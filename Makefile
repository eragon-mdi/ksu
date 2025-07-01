rebuild:
	docker compose down
	docker compose build
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

#clean:
#	sudo docker system prune -a --volumes