FROM alpine:3.14.3

WORKDIR /opt

RUN apk update && apk add bash

COPY ctcgrpc ctcgrpc

ENTRYPOINT ["./ctcgrpc"]
