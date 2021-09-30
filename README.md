# Golang SMS API
## Usage
### Preparations
Edit `config.yaml`

### Installation
`$ go build`

### Running
`$ sudo ./go-sms-api` (sudo is required)

## Development
### Dependencies
```
$ export GO111MODULE=off

$ go get github.com/grpc-ecosystem/grpc-gateway

$ go get github.com/envoyproxy/protoc-gen-validate

$ cd ~/go/src/github.com/grpc-ecosystem/grpc-gateway
$ mkdir third_party && cd third_party
$ git clone https://github.com/googleapis/googleapis.git
```