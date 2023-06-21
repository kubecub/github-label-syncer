FROM golang:1.18 AS build

WORKDIR /go/src/app
COPY . /go/src/app
RUN go get -d -v ./...
# RUN go build -o /go/bin/app cmd/exporter/main.go
RUN go build -o /go/bin/app cmd/syncer/main.go

FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
CMD ["/app"]
