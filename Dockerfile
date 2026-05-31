FROM --platform=$BUILDPLATFORM golang:1.26.1 AS build

ARG TARGETOS=linux
ARG TARGETARCH=arm64

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY src ./src

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /demo-backend ./src

FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /demo-backend /demo-backend

ENTRYPOINT ["/demo-backend"]
