#!/bin/bash

# Остановка контейнера
docker compose stop auth-prod || true

# Удаление старых контейнеров
docker rm $(docker ps -aq) || true

# Удаление старых образов
docker rmi $(docker images -aq) || true

# Запуск контейнеров
docker compose up -d migrator-prod auth-prod  || true
