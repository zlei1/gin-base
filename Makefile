build:
	@go build -v .

dc-start:
	@docker-compose up -d

dc-status:
	@docker-compose ps

dc-clean:
	@echo "Clean Docker images..."
	@docker ps -aqf status=exited | xargs docker rm && docker images -qf dangling=true | xargs docker rmi
