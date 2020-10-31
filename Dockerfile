FROM golang:1.15

WORKDIR /go/src/app
COPY . .

RUN ["go", "test", "github.com/eepnv/example_adapter_unittest/layers"]