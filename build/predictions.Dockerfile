FROM golang:latest as builder
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o predictions -i cmd/predictions/predictions.go


FROM alpine
WORKDIR /app

COPY --from=builder app/predictions .
#COPY --from=builder app/ .

CMD sleep 15 && ./predictions