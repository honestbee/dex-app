# -- Go Builder Image --
FROM golang:1.8-alpine AS builder

RUN apk add --no-cache git
COPY . /go/src/dex-app
WORKDIR /go/src/dex-app

RUN set -ex \
    && go get -v \
    && go build -v -o "/dex-app"

# -- dex-app Image --
FROM alpine:3.6
RUN set -ex \
    && apk add --no-cache bash ca-certificates
COPY entrypoint.sh /entrypoint.sh

COPY --from=builder /dex-app /bin/dex-app
COPY --from=builder /go/src/dex-app/templates/ /bin/templates/
COPY --from=builder /go/src/dex-app/assets/ /bin/assets/
ENTRYPOINT [ "/entrypoint.sh" ]
