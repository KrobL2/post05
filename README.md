# BarberSchedule

**BarberSchedule** - это полноценный RESTful server, это сервис предназначенный для работы с расписанием для добавления и удаления записей в парикмахерской.

## О репозитории

| Папка      | Описание                       |
|------------|--------------------------------|
| cmd        | main.go              | Go      |
| .github    | GitHub Actions       | yaml    |


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
1) Название Image: barber_schedule
2) Название контейнера: _


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
docker run -p 80:8080 имя_образа
```
Где 80 - это порт на хосте, а 8080 - это порт внутри контейнера.


```bash
docker exec -it xbarber-app /bin/sh
```





====

# build both images
docker buildx build --platform linux/arm64,linux/amd64 .
# load just one platform
docker buildx build --load --platform linux/amd64 -t my-image-tag .
# optionally load another platform with a different tag
docker buildx build --load --platform linux/arm64 -t my-image-tag:arm64 .


# push both platforms as one image manifest list
docker buildx build --push --platform linux/arm64,linux/amd64 -t docker.palantir.build/publish:tag .