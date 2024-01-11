run:
	go run cmd/app/app.go

swag-init:
	swag init -g cmd/app/app.go

## Docker Compose
docker-compose:
	docker-compose up -d

docker-compose-restart:
	docker-compose down
	docker-compose up -d

docker-compose-delete:
	docker-compose down
