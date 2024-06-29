FROM golang:1.22-alpine3.19 AS build

LABEL authors="Benyamin Mahmoudyan"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

COPY . .

RUN go build -o /app/hellofresh


FROM  alpine:3.19 AS RUN

COPY --from=build /app/hellofresh /bin/hellofresh

EXPOSE 1986

CMD ["/bin/hellofresh"]
