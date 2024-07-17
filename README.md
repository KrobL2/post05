# RESTful server

В проекте использованы:
1. Docker
2. Docker-compose
3. Github runners for a Go project



====
The server is set to listen on PORT ":1234" by default, which is correct.

However, when running in a Docker container, "localhost" refers to the container itself, not your host machine.

To fix this, you need to ensure that you're exposing the container's port 1234 and mapping it to a port on your host machine when running the Docker container.

Try running your Docker container with the following command:

docker run -p 1234:1234 your-image-name



This maps the container's port 1234 to port 1234 on your host machine.

After doing this, you should be able to access the server using:

curl http://localhost:1234/time
