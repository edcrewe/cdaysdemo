# Demo JSON Transcoding with Envoy

This project demonstrates how to use Envoy's [gRPC-JSON transcoder](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter) to expose gRPC services as RESTful JSON APIs.

## Features

- gRPC service definition
- Envoy proxy configuration for JSON transcoding
- Example REST and gRPC requests

Everything is run locally (rather than via a docker compose) - to make it clear how everything is installed and generated. Also to make it easy to edit / break things, tail the servers, and see what is going on.

You will need a laptop with Go installed.

In the real world this would all be installed in K8s and probably using Istio mesh rather than a single Envoy server.

## Getting Started

1. **Clone the repository**
    ```sh
    git clone https://github.com/your-org/cdaysdemo.git
    cd cdaysdemo
    ```

2. **Install protoc and libs**

    See https://protobuf.dev/installation/ for your enviorment.
    Then install the generator libs.

    ```sh
    ./install_protoc.sh
    ```

3. **Generate the code user protobuf**

    Assuming you have recent Go installed on your laptop you can run generate.sh

    Run generate creates a new clean generated folder
    Gets the required extras for protobuf, annotations and http, plus the protovalidate
    Runs the protoc commands to generate the Go code and the REST API code

    ```sh
    ./generate.sh
    ```

4. **Run up the gRPC server**

    ```sh
    ./run_go.sh

    Starting gRPC server on :9090
    2026/01/18 16:07:04 gRPC server starting on :9090
    ```

    If you have a gRPC client such as Kreya you can now try out the endpoints

5. **Run the Envoy Service**

    Install envoy on your laptop
    https://www.envoyproxy.io/docs/envoy/latest/start/install

    ```sh
    ./run_envoy.sh

    [2026-01-18 16:08:21.448][24239152][debug][upstream] [source/common/upstream/cluster_manager_impl.cc:1506] membership update for TLS cluster cdaysdemo added 1 removed 0
    [2026-01-18 16:08:21.448][24239121][info][main] [source/server/server.cc:1051] starting main dispatch loop
    [2026-01-18 16:08:26.448][24239121][debug][main] [source/server/server.cc:246] flushing stats
    [2026-01-18 16:08:31.449][24239121][debug][main] [source/server/server.cc:246] flushing stats
    [2026-01-18 16:08:36.451][24239121][debug][main] [source/server/server.cc:246] flushing stats
    [2026-01-18 16:08:41.452][24239121][debug][main] [source/server/server.cc:246] flushing stats
    ```

6. **Test the REST endpoint**

    ```sh
    # Index page:
    curl http://localhost:8888/v1/index.html

    # Create: 
    curl -X POST -d '{"id": 1, "name": "Sprocket"}' http://localhost:8888/v1/widget

    # List:
    curl http://localhost:8888/v1/widget

    # Get Specific:
    curl http://localhost:8888/v1/widget/1

    # Delete:
    curl -X DELETE http://localhost:8888/v1/widget/1

    # Get File
    curl http://localhost:8888/v1/file/small.csv

    # Create a big CSV file over 1Gb to demo
    cmd/server/get_big_csv.sh
    # Stream a Big File
    curl http://localhost:8888/v1/stream/big.csv
    ```

## Resources

- [Envoy gRPC-JSON Transcoder Docs](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter)
- [gRPC Documentation](https://grpc.io/docs/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Protocol Buffers Language Guide](https://developers.google.com/protocol-buffers/docs/overview)