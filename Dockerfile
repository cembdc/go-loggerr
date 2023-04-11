FROM golang:1.20-alpine as build

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

FROM alpine

COPY --from=build /app/bin/app /app

ENTRYPOINT ["/app"]
