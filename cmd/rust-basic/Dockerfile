FROM maxmcd/wasm-servers:runtime as runtime
FROM rust:1.30-slim

RUN rustup target add wasm32-unknown-unknown \
    && cargo install --git https://github.com/alexcrichton/wasm-gc

WORKDIR /go/src/github.com/maxmcd/wasm-servers/cmd/rust-basic/
COPY ./Cargo.toml .
COPY ./Cargo.lock .
RUN mkdir src && touch src/lib.rs && cargo build --release

COPY --from=runtime /go/src/github.com/maxmcd/wasm-servers/cmd/runtime/runtime /usr/local/bin/runtime
COPY ./src ./src

RUN cargo build --target wasm32-unknown-unknown --release \
    && wasm-gc target/wasm32-unknown-unknown/release/rust_basic.wasm

CMD runtime target/wasm32-unknown-unknown/release/rust_basic.wasm
