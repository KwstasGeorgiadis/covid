FROM golang:alpine

RUN export env19="./config/covid.docker.json"

RUN mkdir /var/log/covid
RUN touch /var/log/covid/app.ndjson
RUN chmod -R 0777 /var/log/covid/app.ndjson
RUN mkdir /app

ADD . /app/
WORKDIR /app

RUN go get github.com/junkd0g/covid

RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]

# Document that the service listens on port 9080.
EXPOSE 9080
