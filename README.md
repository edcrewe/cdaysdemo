# Demo JSON Transcoding with Envoy

This project demonstrates how to use Envoy's [gRPC-JSON transcoder](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter) to expose gRPC services as RESTful JSON APIs.

## Features

- gRPC service definition
- Envoy proxy configuration for JSON transcoding
- Example REST and gRPC requests

## Getting Started

1. **Clone the repository**
    ```sh
    git clone https://github.com/your-org/cdaysdemo.git
    cd cdaysdemo
    ```

2. **Generate the code user protobuf**

    Assuming you have recent Go installed on your laptop you can run generate.sh

    Run generate creates a new clean generated folder
    Gets the required extras for protobuf, annotations and http, plus the protovalidate
    Runs the protoc commands to generate the Go code and the REST API code

    ```sh
    ./generate.sh
    ```

3. **Test the REST endpoint**
    ```sh
    curl -X POST http://localhost:8080/v1/example -d '{"name": "World"}'
    ```

## Resources

- [Envoy gRPC-JSON Transcoder Docs](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter)
- [gRPC Documentation](https://grpc.io/docs/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Protocol Buffers Language Guide](https://developers.google.com/protocol-buffers/docs/overview)