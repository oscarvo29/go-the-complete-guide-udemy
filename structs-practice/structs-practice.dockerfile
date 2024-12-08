FROM alpine:latest

RUN mkdir /app

COPY structApp /app

CMD ["/app/structApp"]