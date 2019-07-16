FROM golang:1.12-alpine

WORKDIR /src
ENV GOMODULE=on
ENV RUNNER_IGNORED=nuxt
EXPOSE 8080

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh

CMD ["fresh"]
