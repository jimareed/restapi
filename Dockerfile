FROM golang:1.10-alpine AS builder
RUN apk --update add git

WORKDIR /go/src/github.com/jimareed/restapi/
COPY . /go/src/github.com/jimareed/restapi/
RUN rm -f goapp
RUN go get github.com/gorilla/mux
RUN go build -o ./goapp ./main.go

FROM alpine:3.6

EXPOSE 8080
COPY --from=builder /go/src/github.com/jimareed/restapi/goapp /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/goapp && chmod +x /usr/local/bin/goapp
USER nobody
CMD goapp
