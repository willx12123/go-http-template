FROM debian:latest

RUN apt-get update && apt-get install -y ca-certificates

COPY ./build /app

WORKDIR /app

EXPOSE 8080

CMD ["bash", "./bootstrap.sh"]
