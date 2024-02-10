FROM golang:alpine3.16
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -o /docker-gs-ping

EXPOSE 8080
EXPOSE 10002
EXPOSE 10003

CMD [ "/docker-gs-ping", "serve" ]
