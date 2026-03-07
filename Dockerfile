FROM golang:1.26.1

WORKDIR /go/src/github.com/kemper0530/demo-backend

# 依存解決レイヤーを先に分離してキャッシュを効かせる
COPY go.mod go.sum ./
RUN go mod download

COPY src ./src

RUN GOOS=linux GOARCH=arm64 go build -o demo-backend ./src

ENTRYPOINT ["/go/src/github.com/kemper0530/demo-backend/demo-backend"]
