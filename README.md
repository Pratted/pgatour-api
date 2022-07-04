# PGA Tour API
This repo consists of several [OpenAPI](https://github.com/OAI/OpenAPI-Specification) definitions and HTTP clients for the [pgatour.com](https://pgatour.com) API.

## Background
The PGA Tour website exposes several different APIs and servers that provide interesting golf data. There are various uses for the golf data, but it is often challenging to aggregate the data together because it lives behind different APIs. This purpose of this repo is to aggregate the useful PGA Tour APIs into the same location with consistent OpenAPI definitions so HTTP Clients can be generated. There are numerous tools that can be used to generate client/servers from an OpenAPI specification, but this project only uses the [oapi-codegen](https://github.com/deepmap/oapi-codegen) tool for Go.

## Client Generation
PGA Tour HTTP clients can be generated using definitions in the `./openapi` directory. Clients are language agonstic meaning a `go`, `java`. `python`, `nodejs`, etc client can be generated from the same OpenAPI spec. The PGA Tour HTTP clients live in `client/` subdirectory and are further partitioned by language. For example, the `client/go` subdirectory contains a Go implementation of several PGA Tour HTTP clients. The steps to generate a client a largely dependent on the language and generation tool, so refer to the languages README file for more information (e.g `client/go/README.md`)

## Updating the Response Schemas
Peridiocally the PGA Tour will update the schema of their json responses. As far as I'm aware, the schema isn't publicly defined anywhere so several steps need to be done to update it.

The [transform.tools](https://transform.tools/json-to-json-schema) has several useful transformation tools that can take a json document and convert it into a JSON schema. To generate the OpenAPI schema in YAML format use the following steps.

1. Convert JSON repsonse into JSON Schema
2. Convert JSON schema from 2 to OpenAPI Schema
3. Convert JSON OpenAPI Schema from 3 to YAML
4. Paste the YAML into corresponding `openapi/components/<name>`.yaml file

###### Ideally these steps would be automated, but I'm lazy and don't want to pull into a bunch of NodeJS dependencies to do that yet so for now use transform.tools. Furthermore, this could be fully automated in a CI/CD pipeline where some bot periodically fetches the JSON, runs it through the converter, checks for schema changes and if it has makes a PR with the relevant changes.

Whenever the OpenAPI definition or schemas are updated the HTTP Clients should be regenerated to keep them in sync.
