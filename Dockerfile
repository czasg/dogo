# first stage build
FROM golang:1.21 as firstStageBuilder

WORKDIR /workspace
COPY . /workspace

ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0

RUN go mod vendor
RUN go build -o /go/bin/app /workspace/main.go

# second stage build
FROM ubuntu:latest

COPY --from=firstStageBuilder /go/bin/app /usr/local/bin/app

ENTRYPOINT ["app"]
CMD ["webserver"]
