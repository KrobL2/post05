# BarberSchedule

**BarberSchedule** - это полноценный RESTful server, это сервис предназначенный для работы с расписанием для добавления и удаления записей в парикмахерской.

## О репозитории

| Папка   | Описание       |
| ------- | -------------- | ---- |
| cmd     | main.go        | Go   |
| .github | GitHub Actions | yaml |

## В проекте использованы:

1. Docker
2. Docker-compose
3. Github runners
4. Gorilla/mux

## Про Github actions:

Создаем файл с кодом

name: Go + PostgreSQL

on: [push]

jobs:
build:
runs-on: ubuntu-latest

    steps:
    # Checks-out your repository under $GITHUB_WORKSPACE, so your job can access it
    - uses: actions/checkout@v4
    - uses: actions/setup-go@v4
      with:
        go-version: '1.21'
        check-latest: true

    - name: Publish Docker Image
      env:
         USERNAME: ${{ secrets.USERNAME }}
         PASSWORD: ${{ secrets.PASSWORD }}
         IMAGE_NAME: barber_schedule
      run: |
        docker images
        docker build -t "$IMAGE_NAME" .
        docker images
        echo "$PASSWORD" | docker login --username "$USERNAME" --password-stdin
        docker tag "${IMAGE_NAME}" "$USERNAME/${IMAGE_NAME}:latest"
        docker push "$USERNAME/${IMAGE_NAME}:latest"

## Про Docker:

1. Название Image: barber_schedule
2. Название контейнера: \_

## Необходимые шаги для развертывания докеризированного приложения, сохраненного в репозитории Git.

Шаги, необходимые для развертывания приложения зависят от окружения, основной процесс развертывания будет таким:

1. Сборка приложения с использованием Docker build в каталоге с кодом приложения
2. Тестирование образа
3. Выгрузка образа в Registry
4. Уведомление удаленного сервера приложений, что он может скачать образ из Registry и запустить его
5. Перестановка порта в прокси HTTP(S)
6. Остановка старого контейнера

## Команды

```bash
docker compose build
```

```bash
docker compose up
```

```bash
docker exec -it xbarber-app /bin/sh
```


# rest-db
Working with a database

Creating database schema:

```bash
psql -U mtsouk postgres -h 0.0.0.0 -p 5436 < create_db.sql
```



## Обращение к БД

curl -s -X GET -H 'Content-Type: application/json' -d '{"username": "admin", "password" : "admin"}' localhost:8080/getall


curl -X GET -H 'Content-Type: application/json' -d '{"username": "admin", "password" : "admin"}' localhost:8080/logged


curl -X POST -H 'Content-Type: application/json' -d '[{"username": "admin", "password" : "admin", "admin":1}, {"username": "packt", "password" : "admin", "admin":0} ]'  localhost:8080/add


curl -X POST -H 'Content-Type: application/json' -d '{"username": "packt", "password" : "admin"}' localhost:8080/login
