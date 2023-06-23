FROM golang:1.18-alpine

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

ENV CGO_ENABLED=0
ENV ROOT=/go/src/app
WORKDIR ${ROOT}

ENV GO111MODULE=on

COPY . .
EXPOSE 8000

RUN go mod init docker-echo-template && go mod tidy
RUN apk upgrade --update && apk add bash && apk --no-cache add git
RUN go install github.com/cosmtrek/air@v1.44.0
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.0

CMD ["air", "-c", ".air.toml"]