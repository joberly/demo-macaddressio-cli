FROM --platform=$BUILDPLATFORM golang:1.19.3 AS builder

COPY ./ /go/src/demo-macaddressio-cli/

ARG TARGETOS
ARG TARGETARCH

RUN set -e ; \
    cd /go/src/demo-macaddressio-cli/cmd/macaddrio ; \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build . ;

FROM alpine:3.17.0

COPY --from=builder /go/src/demo-macaddressio-cli/cmd/macaddrio/macaddrio /macaddrio

ENTRYPOINT ["/macaddrio"]
