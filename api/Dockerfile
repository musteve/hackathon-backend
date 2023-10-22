FROM golang:1.21-bullseye as builder

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -v -o server

EXPOSE 8080

CMD [ "/app/server" ]