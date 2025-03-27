# ramock
REST API mock service

Simple application for temporary mocking of REST API service that is either not available or not implemented yet.

*ramock* listens for API calls on selected port and returns pre-defined (via config file) http status codes and response body (if given).

If *RAMOCK_STRICT_PATHS* variable is set to *false* (or just not set at all) then *ramock* accepts calls on any endpoint returning HTTP OK (200). For HTTP methods that expects to return body in response (i.e. GET) content of request body is returned.

*ramock* always provide at least */health* endpoint that accepts GET method and returns HTTP OK. If for some reason another response from */health* is needed (i.e. to test orchestrator behavior) then endpoint has to be overwritten in config file.

## Environment variables

| Variable name | Required (Y/N) | Default value | Description |
| --- | --- | --- | --- |
| RAMOCK_LISTEN_PORT | N | 8008 | Application listening port |
| RAMOCK_STRICT_PATHS | N | false | Accept only calls to defined endpoints |
| RAMOCK_ENDPOINTS_FILE | N | endpoints.yaml | Path to YAML file with endpoints definition |

## Endpoints configuration file format

Endpoints configuration for *ramock* should be provided in file with YAML syntax. Current version (0.1) accepts following format:

```yaml
ramockVersion: "0.1"

endpoints:
  - path: "/name1"
    code: 200
    method: "GET"
    body: '{"fieldx":"valueX","fieldY":"valueY"}'
    contentType: "application/json"
  - path: "/name2"
    code: 201
    method: "POST"
    response: '{"fieldA": "value@"}'
    contentType: "application/json"
```

## Running as a Docker/Podman container

### Build by yourself

*ramock* can be build via simple `go build -o ramock ./cmd/ramock/main.go` or using provided *Makefile* via `make ramock`.

When building with *Makefile* output file and endpoints definition example (`endpoints.yaml`) will land in *./build* directory.

It is also possible to directly launch application via `go run ./cmd/ramock/main.go`

### Use image from DockerHub

To try *ramock* without building by yourself or to use it as a part of deployment one can use pre-build container image available on DockerHub at `markamdev/ramock`.

Command below will launch *ramock* with default (sample) configuration.

```bash
<YOUR_CONTAINER_RUNTIME> run -p 8008:8008 --name ramock markamdev/ramock:latest
```

*\<YOUR_CONTAINER_RUNTIME\>* above will be in most cases `podman` or `docker`.

To run *ramock* container with user defined configuration a folder containing *endpoints.yaml* file should be mounted to container's */config* directory:

```bash
# prepare volume directory
mkdir ./ramockConfig
# prepare config file (replace command below with your file fetching)
wget https://raw.githubusercontent.com/markamdev/ramock/refs/heads/master/data/endpoints-example.yaml -O ./ramockConfig/endpoints.yaml
# run container mounting folder as volume
podman run -v ./ramockConfig:/config -p 8008:8008 --name ramock docker.io/markamdev/ramock
```

NOTE: Remember to have folder mounting initialized at `podman machine init` step.

## Author

This project has been created mainly for fun by [*Marcin Kaminski*](mailto:markamdev.84@gmail.com)
