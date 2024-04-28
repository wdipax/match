FROM golang:1.22.2-alpine3.19 AS build
WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -o bot ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /build/bot .
ENTRYPOINT [ "./bot" ]
