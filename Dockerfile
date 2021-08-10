FROM golang:1.15-alpine3.12 as builder

ENV APP_PATH /insights-cache-projects 
COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

RUN go build -o app .

FROM alpine:3.12
WORKDIR /root/
COPY --from=builder /insights-cache-projects/app .

CMD ["./app"]