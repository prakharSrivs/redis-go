FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./

COPY . .
RUN go build -o main ./app/server.go

EXPOSE 6379

CMD [ "./main" ]