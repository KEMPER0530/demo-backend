FROM golang:1.17.2-alpine
MAINTAINER kemper0530

ENV GOPATH /go
ENV PATH=$PATH:$GOPATH/src

# 以下、Docker run 用の設定
ENV PATH=$PATH:$GOPATH/src/github.com/kemper0530/mailform-demo-backend/src
WORKDIR $GOPATH/src/github.com/kemper0530/mailform-demo-backend
COPY  /src $GOPATH/src/github.com/kemper0530/mailform-demo-backend/src

RUN go mod init mailform-demo-backend
RUN go mod tidy

RUN GOOS=linux go build -o mailform-demo-backend ./src

ENTRYPOINT ["/go/src/github.com/kemper0530/mailform-demo-backend/mailform-demo-backend"]
