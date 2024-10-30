#!/bin/bash

# Остановка контейнера
docker compose stop ${AUTH_CONTAINER_NAME} || true

# Удаление старых контейнеров
docker rm $(docker ps -aq) || true

# Удаление старых образов
docker rmi $(docker images -aq) || true

# Запуск контейнеров
docker compose up ${AUTH_CONTAINER_NAME} ${PG_CONTAINER_NAME} || true
