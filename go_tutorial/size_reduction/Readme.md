
# Commands

```bash
# build image
docker build -t size_redduction .

# run image
docker run -d -it -p 5000:80 -v $(pwd)/logs/:/app/logs --name size_redduction size_redduction:latest

# get container ipAddress
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' size_redduction

```
