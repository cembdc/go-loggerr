FROM golang:1.20-alpine as build

RUN apk --no-cache add make

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN make

FROM alpine

COPY --from=build /app/bin/app /app

ENTRYPOINT ["/app"]
