# ramock
REST API mock service

Simple application for temporary mocking of REST API service that is either not available or not implemented yet.

_ramock_ listens for API calls on selected port and returns pre-defined (via config file) http status codes and response body (if given).

If _RAMOCK_STRICT_PATHS_ variable is set to _false_ (or just not set at all) then _ramock_ accepts calls on any endpoint returning HTTP OK (200). For HTTP methods that expects to return body in response (i.e. GET) content of request body is returned.

_ramock_ always provide at least _/health_ endpoint that accepts GET method and returns HTTP OK. If for some reason another response from _/health_ is needed (i.e. to test orchestrator behavior) then endpoint has to be overwritten in config file.

## Environment variables

| Variable name | Required (Y/N) | Default value | Description |
| --- | --- | --- | --- |
| RAMOCK_LISTEN_PORT | Y | | Application listening port |
| RAMOCK_STRICT_PATHS | N| false | Accept only calls to defined endpoints |

## API configuration file format

...