FROM jfloff/alpine-python:2.7-slim
LABEL maintainer="547617251lrtz@gmail.com"
RUN pip install flask
WORKDIR /app
RUN ["mkdir", "logs"]
COPY . .
ENTRYPOINT ["python", "app.py"]