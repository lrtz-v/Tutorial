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

# expose暴露容器给link到当前容器的容器，或者暴露给同一个networks的容器
EXPOSE 80
ENTRYPOINT ["python", "app.py"]

# VOLUME /app
