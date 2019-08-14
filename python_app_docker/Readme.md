
# Commands

```bash
# build image
docker build -t python_docker_app_test .

# run image
docker run -d -it -p 5000:80 -v $(pwd)/logs/:/app/logs --name python_docker_test4 python_docker_app_test:latest

# get container ipAddress
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ecstatic_sammet

```
