FROM python:2.7.10
LABEL maintainer="547617251lrtz@gmail.com"

# RUN apk update && apk upgrade && apk add bash

# Use ARG values to set a default value for an ENV variable from the command line when build the image
ARG APP=Tutorials
# ENV variable persists through container run time
ENV APP=python_app_docker_test
ENV LANG=en_US.UTF-8 LC_ALL=en_US.UTF-8

RUN pip install flask

WORKDIR /app
RUN ["mkdir", "logs"]
COPY . .

# Favor ENTRYPOINT when you need to run the same command every time.
# Favor ENTRYPOINT when a container will be used as an executable program.
# Favor CMD when you need to provide extra default arguments that could be overwritten from the command line.
ENTRYPOINT ["python", "app.py"]
# CMD ["python", "./app.py"]

# expose暴露容器给link到当前容器的容器，或者暴露给同一个networks的容器
# EXPOSE 5000

# VOLUME /app
