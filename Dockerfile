FROM ubuntu:latest

WORKDIR /app/watchdog

COPY . .

ENTRYPOINT [ "./main" ]