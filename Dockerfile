FROM --platform=$BUILDPLATFORM golang:1.19.3 AS builder

COPY ./ /go/src/demo-macaddressio-cli/

ARG TARGETOS
ARG TARGETARCH

RUN set -e ; \
    cd /go/src/demo-macaddressio-cli/cmd/macaddrio ; \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build . ;

FROM debian:buster-slim

RUN set -e ; \
    apt update ; \
    apt install -y ca-certificates ; \
    apt clean ;

COPY --from=builder /go/src/demo-macaddressio-cli/cmd/macaddrio/macaddrio /macaddrio

ENTRYPOINT ["/macaddrio"]
