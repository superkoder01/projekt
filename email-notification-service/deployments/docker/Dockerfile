FROM alpine:3

RUN apk update && apk upgrade && apk add bash libc6-compat tzdata
RUN mkdir -p /opt/c4e

ENV TZ=Europe/Warsaw

WORKDIR /opt/c4e

COPY ./email_service .
COPY ./email-notification-service.yaml .

CMD ["./email_service"]