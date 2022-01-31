FROM golang:1.17-buster
WORKDIR /build
COPY . .
RUN go build -o ./target/color ./cmd/...
RUN mv /build/target/color /usr/bin/color
RUN chmod +x /usr/bin/color
ENV COLOR_HTTP_ADDRESS=0.0.0.0:8080
ENTRYPOINT ["/usr/bin/color"]
