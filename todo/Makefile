DOCKER=docker
COMPOSE=$(DOCKER) compose
EXEC=$(COMPOSE) exec
EXECD=$(COMPOSE) exec -d
BUILD=$(COMPOSE) build
UP=$(COMPOSE) up -d
LOGS=$(COMPOSE) logs
STOP=$(COMPOSE) stop
RM=$(COMPOSE) rm
DOWN=$(COMPOSE) down
BACKEND=$(EXEC) backend
FRONTEND=$(EXEC) frontend

DB=$(EXEC) db

MYSQL_USER=user
MYSQL_PASSWORD=password
MYSQL_DATABASE=todo

all: docker/up

docker/build: ## docker build
	$(BUILD)

docker/up: ## docker up
	$(UP)

docker/logs: ## docker logs
	$(LOGS)

docker/stop: ## docker stop
	$(STOP)

docker/rm: ## docker clean
	$(RM)

docker/down: ## docker down
	$(DOWN) -v

docker/volume/prune: ### docker volume prune
	$(DOCKER) volume prune

backend/bash: ## backend container bash
	$(BACKEND) bash

frontend/bash: ## frontend container bash
	$(FRONTEND) bash

db/bash: ## db(postgres) container bash
	$(DB) bash

mysql: ## db(mysql) access
	$(DB) bash -c "mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
