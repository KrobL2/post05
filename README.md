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

====
The server is set to listen on PORT ":1234" by default, which is correct.

However, when running in a Docker container, "localhost" refers to the container itself, not your host machine.

To fix this, you need to ensure that you're exposing the container's port 1234 and mapping it to a port on your host machine when running the Docker container.

Try running your Docker container with the following command:

docker run -p 1234:1234 your-image-name



This maps the container's port 1234 to port 1234 on your host machine.

After doing this, you should be able to access the server using:

curl http://localhost:1234/time



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
docker run -p 80:8080 имя_образа
```
Где 80 - это порт на хосте, а 8080 - это порт внутри контейнера.