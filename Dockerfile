FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum main.go ./
COPY config ./config
COPY data ./data
COPY handler ./handler

RUN ls -al
RUN go mod download
RUN go mod tidy

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main .

ENV PROFILE=prod \
    ALLOWED_ORIGINS="http://localhost:3000,http://133.186.228.94:31030"

ENTRYPOINT ["/main"]
