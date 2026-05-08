# syntax=docker/dockerfile:1

# BASE
ARG GO_VERSION=1.25.4
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-bookworm AS base
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# BUILD
FROM --platform=$BUILDPLATFORM base AS build

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/appname cmd/main.go

# VAT
FROM --platform=$BUILDPLATFORM base AS vet

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go vet

# TEST
FROM --platform=$BUILDPLATFORM base AS test

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go test # TODO divide unitt&dbt

# RUN
FROM debian:bookworm-slim AS final

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

ENV PORT 8000
ENV DB_DRIVER ""
ENV DB_NAME ""
ENV DB_HOST ""
ENV DB_PORT ""
ENV DB_USER ""
ENV DB_PASSWORD ""

COPY --from=build /bin/server /bin/

EXPOSE 8000

ENTRYPOINT [ "/bin/appname" ]
