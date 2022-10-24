FROM alpine:3.6 as alpine

RUN apk add -U --no-cache ca-certificates tzdata

ARG TZ=Asia/Shanghai
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV ENVIRONMENT TEST

ADD dist/recative-content-auth-service /app/

EXPOSE 80

WORKDIR /app

ENTRYPOINT ["/app/recative-content-auth-service"]