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


FROM python:3.8-slim

WORKDIR /app

COPY --from=builder app/main .
COPY --from=builder app/config.json .
COPY --from=builder app/internal/pkg/image/analyze.py .
COPY --from=builder app/internal/pkg/image/data.csv .

RUN pip install torch==1.5.0+cpu torchvision==0.6.0+cpu -f https://download.pytorch.org/whl/torch_stable.html

CMD sleep 15 && ./main