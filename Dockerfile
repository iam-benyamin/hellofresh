FROM golang:1.22-alpine3.19 AS build

LABEL authors="Benyamin Mahmoudyan"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

COPY . .

ARG SERVICE

RUN go build -o /app/bin/${SERVICE} /app/cmd/${SERVICE}/main.go

FROM alpine:3.19 AS run

ARG SERVICE

COPY config.yaml /bin/config.yaml

COPY --from=build /app/bin/${SERVICE} /bin/
