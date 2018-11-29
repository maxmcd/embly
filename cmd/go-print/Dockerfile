FROM maxmcd/wasm-servers:go

RUN mkdir -p /go/src/github.com/perlin-network/life/ \
    && cd /go/src/github.com/perlin-network/life \
    && git clone https://github.com/perlin-network/life.git . \
    && GO111MODULE=on go mod vendor 

WORKDIR /go/src/github.com/maxmcd/wasm-servers/cmd/go-print

COPY ./go-print .
COPY ./runtime ../runtime
RUN cd ../runtime && go install

RUN GOOS=js GOARCH=wasm go build -o go-print.wasm 

CMD runtime ./go-print.wasm
