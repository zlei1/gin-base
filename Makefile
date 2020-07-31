start:
	@docker-compose up -d

status:
	@docker-compose ps

clean:
	@echo "Clean Docker images..."
	@docker ps -aqf status=exited | xargs docker rm && docker images -qf dangling=true | xargs docker rmi
