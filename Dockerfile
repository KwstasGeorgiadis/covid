FROM golang:alpine

RUN mkdir /var/log/covid
RUN touch /var/log/covid/app.ndjson
RUN chmod -R 0777 /var/log/covid/app.ndjson
RUN mkdir /app

ADD . /app/
WORKDIR /app
RUN apk add git

RUN git clone https://github.com/junkd0g/covid.git



RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]

# Document that the service listens on port 9080.
EXPOSE 9080
