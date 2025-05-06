# [Arg variable]
ARG WORK_DIR="/go/src/app"

# build (for development)
FROM golang:1.24-alpine AS builder
ARG WORK_DIR
WORKDIR ${WORK_DIR}

COPY . .
RUN apk upgrade --update && apk add --no-cache gcc musl-dev make git curl
RUN make build

# release
FROM alpine:latest AS app
ARG WORK_DIR
COPY --from=builder ${WORK_DIR}/bin/app /usr/local/bin/app
ENTRYPOINT ["app"]
