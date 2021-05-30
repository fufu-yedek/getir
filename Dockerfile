FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o getir-build .

WORKDIR /dist

RUN cp /build/getir-build .
RUN cp /build/static/config.json .

ENV GETIR_CONFIG_PATH=/dist/config.json

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/getir-build"]