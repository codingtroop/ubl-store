FROM golang:1.17-alpine as build

WORKDIR /app

RUN apk  --no-cache --update upgrade && apk --no-cache add gcc ca-certificates musl-dev

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN swag init -g ./main.go \
    && go build -o /ubl-store github.com/codingtroop/ubl-store


FROM alpine:3.14

WORKDIR /app
COPY --from=build /ubl-store .
COPY --from=build /app/config.yml .

EXPOSE 80

CMD ["./ubl-store"]