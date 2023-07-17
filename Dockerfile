# Copyright © 2023 KubeCub open source community. All rights reserved.
# Licensed under the MIT License (the "License");
# you may not use this file except in compliance with the License.

FROM golang AS build
ENV GOPROXY=https://goproxy.cn

WORKDIR /go/src/app
COPY . /go/src/app
# RUN go build -o /go/bin/app cmd/exporter/main.go
RUN go build -o /go/bin/app cmd/syncer/main.go

FROM alpine
COPY --from=build /go/bin/app /
CMD ["/app"]
