# Copyright © 2023 KubeCub open source community. All rights reserved.
# Licensed under the MIT License (the "License");
# you may not use this file except in compliance with the License.

FROM golang:1.20 AS build

WORKDIR /go/src/app
COPY . /go/src/app
RUN go get -d -v ./...
# RUN go build -o /go/bin/app cmd/exporter/main.go
RUN go build -o /go/bin/app cmd/syncer/main.go

FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
CMD ["/app"]
