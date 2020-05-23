FROM golang:latest as builder
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main -i cmd/main/main.go

FROM alpine
WORKDIR /app

COPY --from=builder app/main .
COPY --from=builder app/config.json .

CMD sleep 10 && ./main