FROM golang:latest

ENV LANG en_US.UTF-8

WORKDIR /app

COPY . .

RUN go get -t -d -v ./...
RUN go build -o /bin/app *.go

CMD [ "app" ]

ENV PORT 8080
EXPOSE 8080
