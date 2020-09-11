build:
	@go build -v .

docker-start:
	@docker-compose up -d

docker-status:
	@docker-compose ps

docker-clean:
	@echo "Clean Docker images..."
	@docker ps -aqf status=exited | xargs docker rm && docker images -qf dangling=true | xargs docker rmi

swag-init:
	@swag init
	@echo "swag init done"
