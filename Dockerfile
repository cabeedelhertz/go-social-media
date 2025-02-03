FROM golang:1.22-alpine as build
WORKDIR /build
RUN apk add --no-cache git
COPY . .

# Environment.
ENV GO111MODULE=on
ENV GOPRIVATE=buf.build/gen/go
ENV CGO_ENABLED=0


# Install and authenticate to BUF to pull remote packages
RUN go install github.com/bufbuild/buf/cmd/buf@v1.25.0
RUN --mount=type=secret,id=BUF_USER \
    --mount=type=secret,id=BUF_TOKEN \
    echo $(cat /run/secrets/BUF_TOKEN) | buf registry login --username $(cat /run/secrets/BUF_USER) --token-stdin

RUN go build -o /social ./cmd/social/...

# Deployment
FROM alpine:3.18 as deploy
COPY --from=build /social /social
EXPOSE 8080
EXPOSE 50051
ENTRYPOINT ["/social"]
CMD ["server"]