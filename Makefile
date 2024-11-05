# Путь к файлам Docker Compose
COMPOSE_PATH = app
COMPOSE_FILES = -f $(COMPOSE_PATH)/docker-compose.yaml -f $(COMPOSE_PATH)/docker-compose.db.yml -f $(COMPOSE_PATH)/docker-compose.queue.yml

# Цели
up:
	docker-compose $(COMPOSE_FILES) up &
	cd frontend && npm run dev

up-build:
	docker-compose $(COMPOSE_FILES) up --build &
	cd frontend && npm run dev

down:
	docker-compose $(COMPOSE_FILES) down

# Удаление неиспользуемых ресурсов
prune:
	docker system prune -f --volumes
