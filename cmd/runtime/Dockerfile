FROM golang:1.11

RUN mkdir -p /go/src/github.com/perlin-network/life/ \
    && cd /go/src/github.com/perlin-network/life \
    && git clone https://github.com/perlin-network/life.git . \
    && GO111MODULE=on go mod vendor 

WORKDIR /go/src/github.com/maxmcd/wasm-servers/cmd/runtime
COPY . .

RUN go build 
