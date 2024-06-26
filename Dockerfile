FROM golang:1.22.2-alpine3.19 AS build
RUN apk add build-base
WORKDIR /build
COPY . .
RUN go mod tidy
ARG LANG="en"
RUN CGO_ENABLED=1 go test -race -tags=${LANG} ./...
RUN go build -tags=${LANG} -trimpath -o bot ./cmd/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=build /build/bot .
ENTRYPOINT [ "./bot" ]
