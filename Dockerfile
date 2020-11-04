FROM docker.io/library/golang:1.15

WORKDIR /go/src/github.com/kubenetic/color-app
ADD . .

RUN GO111MODULE=on CGO_ENABLED=0 go build github.com/kubenetic/color-app

#################################################
FROM docker.io/library/alpine:latest

WORKDIR /app

ADD ./templates/ /app/templates/
COPY --from=0 /go/src/github.com/kubenetic/color-app/color-app .

EXPOSE 8080

ENTRYPOINT ["/app/color-app"]