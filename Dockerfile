FROM golang:1.17.2-alpine
MAINTAINER kemper0530

ENV GOPATH /go
ENV PATH=$PATH:$GOPATH/src

# 以下、Docker run 用の設定
ENV PATH=$PATH:$GOPATH/src/github.com/kemper0530/demo-backend/src
WORKDIR $GOPATH/src/github.com/kemper0530/demo-backend
COPY  /src $GOPATH/src/github.com/kemper0530/demo-backend/src

RUN go mod init demo-backend
RUN go mod tidy

RUN GOOS=linux go build -o demo-backend ./src

ENTRYPOINT ["/go/src/github.com/kemper0530/demo-backend/demo-backend"]
