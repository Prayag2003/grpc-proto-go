# gRPC Proto Go Example

This repository demonstrates how to use Protocol Buffers (`.proto` files) and gRPC in Go.

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Protocol Buffers Compiler (`protoc`)](https://grpc.io/docs/protoc-installation/)

## Installation

Install the Protocol Buffers compiler using Homebrew:

```sh
brew install protoc
```

Verify the installation:

```sh
protoc --version
```

## Usage

1. **Write the `greet.proto` file**

      Create your Protocol Buffers definition in `proto/greet.proto`.

2. **Generate Go code from the proto file**

      Run the following command to generate Go and gRPC code:

      ```sh
      protoc --go_out=. --go-grpc_out=. proto/greet.proto
      ```

      This will generate the necessary Go files for your gRPC service.

## Resources

- [gRPC Go Documentation](https://grpc.io/docs/languages/go/)
- [Protocol Buffers Documentation](https://developers.google.com/protocol-buffers/docs/overview)
